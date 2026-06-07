package dmarket

import (
	"go.uber.org/zap"
)

const (
	DMarketURL = "https://api.dmarket.com/"
)

type DmarketCfg struct {
	Logger  *zap.Logger
	SecKey  *string `json:"secret_key"`
	PrivKey *string `json:"private_key"`
	Proxy   *string `json:"proxy_url"`
}

func NewDmarketClient(cfg DmarketCfg) (*ClientWithResponses, error) {
	opts := []ClientOption{}

	if cfg.SecKey != nil && cfg.PrivKey != nil {
		opt, err := WithAuth(*cfg.PrivKey, *cfg.SecKey)
		if err != nil {
			return nil, err
		}
		opts = append(opts, opt)
	}

	if cfg.Proxy != nil {
		opt, err := WithProxy(*cfg.Proxy)
		if err != nil {
			return nil, err
		}
		opts = append(opts, opt)
	}

	return NewClientWithResponses(DMarketURL, opts...)
}
