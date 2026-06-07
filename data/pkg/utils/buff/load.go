package buff_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/buff"
	buffqueue "github.com/MassimoMarsiglia/dmarket-bot/pkg/models/buff_queue"

	"github.com/gammazero/deque"
	"go.uber.org/zap"
)

type GetListingParams struct {
	MarketHashName string
	Item           buffqueue.Item
}

var ErrNoClients = errors.New("no clients loaded")

var ErrNoMappings = errors.New("no mappings loaded")

var ErrMissingMappingDir = errors.New("no mapping directory provided")

func LoadBuffQueue(logger *zap.Logger, dir string) (*deque.Deque[*GetListingParams], error) {
	logger.Debug("loading buff queue")
	var queue deque.Deque[*GetListingParams]

	if dir == "" {
		return nil, fmt.Errorf("Failed to load buff queue: %w", ErrMissingMappingDir)
	}

	logger.Debug("loading buff queue", zap.String("dir", dir))
	reader, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var buffMappings buffqueue.ItemMapping
	err = json.Unmarshal(data, &buffMappings)
	if err != nil {
		logger.Error("failed to unmarshal mappings",
			zap.Error(err),
		)
		return nil, err
	}

	for key, value := range buffMappings.Items {
		params := GetListingParams{
			MarketHashName: key,
			Item:           value,
		}
		queue.PushBack(&params)
	}
	if queue.Len() == 0 {
		logger.Error("failed to load mappings",
			zap.Error(ErrNoMappings),
			zap.String("mapping dir", dir),
		)
		return nil, ErrNoMappings
	}

	return &queue, nil
}

func LoadBuffAccounts(logger *zap.Logger, accDir *string, cfgs []buff.BuffCfg) (*deque.Deque[*buff.ClientWithResponses], error) {
	logger.Debug("loading buff accounts")
	var clients deque.Deque[*buff.ClientWithResponses]
	if accDir != nil && *accDir != "" {
		dir := *accDir
		logger.Debug("loading buff accounts", zap.String("dir", dir))

		reader, err := os.Open(dir)
		if err != nil {
			return nil, err
		}

		defer reader.Close()
		data, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		var accounts []buff.BuffCfg
		err = json.Unmarshal(data, &accounts)
		if err != nil {
			return nil, err
		}

		clients.SetBaseCap(clients.Len() + len(cfgs))
		for _, acc := range accounts {
			acc.Logger = logger
			acc, err := buff.NewBuffClient(acc)
			if err != nil {
				return nil, err
			}
			clients.PushBack(acc)
		}
	}
	for _, acc := range cfgs {
		acc, err := buff.NewBuffClient(acc)
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
	logger.Debug("loaded clients", zap.Int("clients", clients.Len()))
	return &clients, nil
}
