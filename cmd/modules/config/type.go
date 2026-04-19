package config

import (
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type BuffListingCfg struct {
	AccDir     string `yaml:"acc_path"`
	MappingDir string `yaml:"mapping_path"`
	Enabled    bool   `yaml:"enabled"`
	Conn       *nats.Conn
	Delay      time.Duration
}

type NewListingCfg struct {
	AccDir  string `yaml:"acc_path"`
	Enabled bool   `yaml:"enabled"`
	Conn    *nats.Conn
	Delay   time.Duration
}

type OrderBookCfg struct {
	AccDir  string `yaml:"acc_path"`
	Enabled bool   `yaml:"enabled"`
	ItemDir string `yaml:"item_path"`
	Conn    *nats.Conn
	Delay   time.Duration
}

type Config[T any] struct {
	logger *zap.Logger
	path   string
}

type BotConfig struct {
	Dmarket DmarketConfig `yaml:"Dmarket"`
	Buff    BuffConfig    `yaml:"Buff"`
}

type BuffConfig struct {
	Listing BuffListingCfg `yaml:"Listing"`
}

type DmarketConfig struct {
	NewListing NewListingCfg `yaml:"NewListing"`
	OrderBook  OrderBookCfg  `yaml:"OrderBook"`
}
