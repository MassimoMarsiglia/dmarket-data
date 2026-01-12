package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	DMarketURL = "https://api.dmarket.com/"
)

var ErrFailedToParse = errors.New("failed to parse")

type DmarketCfg struct {
	secKey  *string
	privKey *string
	proxy   *string
}

func NewDmarketClient(cfg DmarketCfg) (*ClientWithResponses, error) {
	opts := []ClientOption{}

	if cfg.secKey != nil && cfg.privKey != nil {
		auth, err := NewDmarketAuth(
			*cfg.privKey,
			*cfg.secKey,
		)
		if err != nil {
			return nil, err
		}
		opts = append(opts, WithRequestEditorFn(auth.Middleware()))
	}

	if cfg.proxy != nil {
		proxyUrl, err := url.Parse(*cfg.proxy)
		if err != nil {
			return nil, fmt.Errorf("%w proxy url %s", ErrFailedToParse, *cfg.proxy)
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
