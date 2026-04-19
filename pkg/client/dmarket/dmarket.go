package dmarket

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

const (
	DMarketURL = "https://api.dmarket.com/"
)

var ErrFailedToParse = errors.New("failed to parse")

type DmarketCfg struct {
	Logger  *zap.Logger
	SecKey  *string `json:"secret_key"`
	PrivKey *string `json:"private_key"`
	Proxy   *string `json:"proxy_url"`
}

func NewDmarketClient(cfg DmarketCfg) (*ClientWithResponses, error) {
	opts := []ClientOption{}

	if cfg.SecKey != nil && cfg.PrivKey != nil {

		privHex := hex.EncodeToString([]byte(*cfg.PrivKey))
		pubHex := hex.EncodeToString([]byte(*cfg.SecKey))

		auth, err := NewDmarketAuth(
			privHex,
			pubHex,
		)
		if err != nil {
			return nil, err
		}
		opts = append(opts, WithRequestEditorFn(auth.Middleware()))
	}

	if cfg.Proxy != nil {
		proxyUrl, err := url.Parse(*cfg.Proxy)
		if err != nil {
			return nil, fmt.Errorf("%w proxy url %s", ErrFailedToParse, *cfg.Proxy)
		}
		opts = append(opts, WithHTTPClient(
			&http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: 10,
					Proxy:               http.ProxyURL(proxyUrl),
				},
			},
		))
	}
	return NewClientWithResponses(DMarketURL, opts...)
}
