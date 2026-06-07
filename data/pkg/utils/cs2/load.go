package cs2

import (
	"encoding/json"
	"io"
	"os"

	CS2Models "github.com/MassimoMarsiglia/dmarket-bot/pkg/models/CS2"
)

func LoadStickers(path string) ([]CS2Models.Sticker, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	data, err := io.ReadAll(reader)

	var stickers []CS2Models.Sticker
	err = json.Unmarshal(data, &stickers)
	if err != nil {
		return nil, err
	}

	return stickers, nil
}

func LoadSkinsNotGrouped(path string) ([]CS2Models.SkinNotGrouped, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var skins []CS2Models.SkinNotGrouped
	err = json.Unmarshal(data, &skins)
	if err != nil {
		return nil, err
	}
	return skins, nil
}

func LoadLookupTable(skinsPath, stickersPath string) (map[string]string, error) {
	skins, err := LoadSkinsNotGrouped(skinsPath)
	if err != nil {
		return nil, err
	}

	stickers, err := LoadStickers(stickersPath)
	if err != nil {
		return nil, err
	}

	entries := make(map[string]string, len(skins)+len(stickers))
	for _, s := range skins {
		entries[s.MarketHashName] = s.SkinId
	}
	for _, s := range stickers {
		entries[s.MarketHashName] = s.Id
	}

	return entries, nil
}
