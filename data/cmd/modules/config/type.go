package config

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/buff/listing"
	newlisting "github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/new_listing"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/orderbook"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type BuffListingCfg struct {
	AccDir     string        `yaml:"acc_path"`
	MappingDir string        `yaml:"mapping_path"`
	Enabled    bool          `yaml:"enabled"`
	Delay      time.Duration `yaml:"delay"`
	Conn       *nats.Conn
	Listing    *listing.Emitter
}

type NewListingCfg struct {
	AccDir     string        `yaml:"acc_path"`
	Enabled    bool          `yaml:"enabled"`
	Delay      time.Duration `yaml:"delay"`
	Conn       *nats.Conn
	NewListing *newlisting.Emitter
}

type OrderBookCfg struct {
	AccDir       string        `yaml:"acc_path"`
	Enabled      bool          `yaml:"enabled"`
	SkinsPath    string        `yaml:"skins_path"`
	StickersPath string        `yaml:"stickers_path"`
	Delay        time.Duration `yaml:"delay"`
	Conn         *nats.Conn
	OrderBook    *orderbook.Emitter
}

type SalesCfg struct {
	AccDir       string        `yaml:"acc_path"`
	Enabled      bool          `yaml:"enabled"`
	SkinsPath    string        `yaml:"skins_path"`
	StickersPath string        `yaml:"stickers_path"`
	Delay        time.Duration `yaml:"delay"`
	Conn         *nats.Conn
}

type Config[T any] struct {
	logger *zap.Logger
	path   string
}

type LookupCfg struct {
	SkinsPath        string `yaml:"skins_path"`
	StickersPath     string `yaml:"stickers_path"`
	AgentsPath       string `yaml:"agents_path"`
	GraffitiPath     string `yaml:"graffiti_path"`
	CratesPath       string `yaml:"crates_path"`
	KeychainsPath    string `yaml:"keychains_path"`
	MusicKitsPath    string `yaml:"music_kits_path"`
	PatchesPath      string `yaml:"patches_path"`
	StickerSlabsPath string `yaml:"sticker_slabs_path"`
	CollectiblesPath string `yaml:"collectibles_path"`
}

type BotConfig struct {
	Dmarket DmarketConfig `yaml:"Dmarket"`
	Buff    BuffConfig    `yaml:"Buff"`
	Lookup  LookupCfg     `yaml:"Lookup"`
}

type BuffConfig struct {
	Listing BuffListingCfg `yaml:"Listing"`
}

type DmarketConfig struct {
	NewListing NewListingCfg `yaml:"NewListing"`
	OrderBook  OrderBookCfg  `yaml:"OrderBook"`
	Sales      SalesCfg      `yaml:"Sales"`
}
