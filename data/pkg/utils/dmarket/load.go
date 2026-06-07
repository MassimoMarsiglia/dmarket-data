package dmarket_utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/cs2"
	"github.com/gammazero/deque"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no clients loaded")

func LoadDmarketAccounts(logger *zap.Logger, accDir *string, cfgs []dmarket.DmarketCfg) (*deque.Deque[*dmarket.ClientWithResponses], error) {
	logger.Debug("loading dmarket accounts")
	var clients deque.Deque[*dmarket.ClientWithResponses]
	if accDir != nil && *accDir != "" {
		dir := *accDir
		logger.Debug("loading dmarket accounts", zap.String("dir", dir))

		reader, err := os.Open(dir)
		if err != nil {
			return nil, err
		}

		defer reader.Close()
		data, err := io.ReadAll(reader)

		var accounts []dmarket.DmarketCfg
		err = json.Unmarshal(data, &accounts)
		if err != nil {
			return nil, err
		}

		clients.SetBaseCap(clients.Len() + len(cfgs))
		for _, acc := range accounts {
			acc.Logger = logger
			acc, err := dmarket.NewDmarketClient(acc)
			if err != nil {
				return nil, err
			}
			clients.PushBack(acc)
		}
	}
	for _, acc := range cfgs {
		acc, err := dmarket.NewDmarketClient(acc)
		if err != nil {
			return nil, err
		}
		clients.PushBack(acc)
	}
	if clients.Len() == 0 {
		logger.Error("failed to load clients",
			zap.Error(ErrNoClients),
			zap.String("account dir", *accDir),
			zap.Int("num cfgs", len(cfgs)))
		return nil, ErrNoClients
	}
	logger.Debug("loaded", zap.Int("clients", clients.Len()))
	return &clients, nil
}

func LoadDmarketOrderBookQueue(logger *zap.Logger, skinsPath, stickersPath string) (*deque.Deque[*dmarket.GetOrderBookParams], error) {
	logger.Debug("Filling order book request queue...")
	var reqs deque.Deque[*dmarket.GetOrderBookParams]

	skins, err := cs2.LoadSkinsNotGrouped(skinsPath)
	if err != nil {
		return nil, err
	}

	stickers, err := cs2.LoadStickers(stickersPath)
	if err != nil {
		return nil, err
	}

	reqs.SetBaseCap(reqs.Len() + len(skins) + len(stickers))

	for i := range skins {
		skin := skins[i]
		reqs.PushBack(&dmarket.GetOrderBookParams{
			Title:  skin.MarketHashName,
			GameId: dmarket.GameIDA8db,
		})
	}

	for i := range stickers {
		sticker := stickers[i]
		reqs.PushBack(&dmarket.GetOrderBookParams{
			Title:  sticker.Name,
			GameId: dmarket.GameIDA8db,
		})
	}

	amount := reqs.Len()
	logger.Debug("Filled order book request queue with", zap.Int("num reqs:", amount))

	return &reqs, nil
}
