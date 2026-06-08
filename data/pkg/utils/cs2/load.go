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

func LoadAgents(path string) ([]CS2Models.Agent, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var agents []CS2Models.Agent
	err = json.Unmarshal(data, &agents)
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func LoadGraffiti(path string) ([]CS2Models.Graffiti, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var graffiti []CS2Models.Graffiti
	err = json.Unmarshal(data, &graffiti)
	if err != nil {
		return nil, err
	}
	return graffiti, nil
}

func LoadCrates(path string) ([]CS2Models.Crate, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var crates []CS2Models.Crate
	err = json.Unmarshal(data, &crates)
	if err != nil {
		return nil, err
	}
	return crates, nil
}

func LoadKeychains(path string) ([]CS2Models.Keychain, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var keychains []CS2Models.Keychain
	err = json.Unmarshal(data, &keychains)
	if err != nil {
		return nil, err
	}
	return keychains, nil
}

func LoadMusicKits(path string) ([]CS2Models.MusicKit, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var musicKits []CS2Models.MusicKit
	err = json.Unmarshal(data, &musicKits)
	if err != nil {
		return nil, err
	}
	return musicKits, nil
}

func LoadPatches(path string) ([]CS2Models.Patch, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var patches []CS2Models.Patch
	err = json.Unmarshal(data, &patches)
	if err != nil {
		return nil, err
	}
	return patches, nil
}

func LoadStickerSlabs(path string) ([]CS2Models.StickerSlab, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var stickerSlabs []CS2Models.StickerSlab
	err = json.Unmarshal(data, &stickerSlabs)
	if err != nil {
		return nil, err
	}
	return stickerSlabs, nil
}

func LoadCollectibles(path string) ([]CS2Models.Collectible, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)

	var collectibles []CS2Models.Collectible
	err = json.Unmarshal(data, &collectibles)
	if err != nil {
		return nil, err
	}
	return collectibles, nil
}

func getMarketHashName(name string, mhn *string) string {
	if mhn != nil && *mhn != "" {
		return *mhn
	}
	return name
}

func LoadLookupTable(skinsPath, stickersPath, agentsPath, graffitiPath, cratesPath, keychainsPath, musicKitsPath, patchesPath, stickerSlabsPath, collectiblesPath string) (map[string]string, error) {
	skins, err := LoadSkinsNotGrouped(skinsPath)
	if err != nil {
		return nil, err
	}

	stickers, err := LoadStickers(stickersPath)
	if err != nil {
		return nil, err
	}

	agents, err := LoadAgents(agentsPath)
	if err != nil {
		return nil, err
	}

	graffiti, err := LoadGraffiti(graffitiPath)
	if err != nil {
		return nil, err
	}

	crates, err := LoadCrates(cratesPath)
	if err != nil {
		return nil, err
	}

	keychains, err := LoadKeychains(keychainsPath)
	if err != nil {
		return nil, err
	}

	musicKits, err := LoadMusicKits(musicKitsPath)
	if err != nil {
		return nil, err
	}

	patches, err := LoadPatches(patchesPath)
	if err != nil {
		return nil, err
	}

	stickerSlabs, err := LoadStickerSlabs(stickerSlabsPath)
	if err != nil {
		return nil, err
	}

	collectibles, err := LoadCollectibles(collectiblesPath)
	if err != nil {
		return nil, err
	}

	entries := make(map[string]string, len(skins)+len(stickers)+len(agents)+len(graffiti)+len(crates)+len(keychains)+len(musicKits)+len(patches)+len(stickerSlabs)+len(collectibles))

	for _, s := range skins {
		entries[s.MarketHashName] = s.SkinId
	}
	for _, s := range stickers {
		entries[s.MarketHashName] = s.Id
	}
	for _, a := range agents {
		entries[a.MarketHashName] = a.Id
	}
	for _, g := range graffiti {
		entries[g.MarketHashName] = g.Id
	}
	for _, c := range crates {
		entries[c.MarketHashName] = c.Id
	}
	for _, k := range keychains {
		entries[k.MarketHashName] = k.Id
	}
	for _, m := range musicKits {
		entries[getMarketHashName(m.Name, m.MarketHashName)] = m.Id
	}
	for _, p := range patches {
		entries[p.MarketHashName] = p.Id
	}
	for _, s := range stickerSlabs {
		entries[getMarketHashName(s.Name, s.MarketHashName)] = s.Id
	}
	for _, c := range collectibles {
		entries[getMarketHashName(c.Name, c.MarketHashName)] = c.Id
	}

	return entries, nil
}
