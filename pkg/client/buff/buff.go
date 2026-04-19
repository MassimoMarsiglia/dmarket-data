package buff

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

const (
	BuffURL = "https://buff.163.com/"
)

var ErrFailedToParse = errors.New("failed to parse")

type BuffCfg struct {
	Logger *zap.Logger
	Proxy  *string `json:"proxy_url"`
}

func BuffHeaders() func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("User-Agent", "curl/8.0.1")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Referer", "https://buff.163.com/market/")
		return nil
	}
}

func NewBuffClient(cfg BuffCfg) (*ClientWithResponses, error) {
	opts := []ClientOption{}

	if cfg.Proxy != nil {
		proxyUrl, err := url.Parse(*cfg.Proxy)
		if err != nil {
			return nil, fmt.Errorf("%w proxy url %s", ErrFailedToParse, *cfg.Proxy)
		}
		_ = proxyUrl
		opts = append(opts, WithHTTPClient(
			&http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: 10,
					Proxy:               http.ProxyURL(proxyUrl),
				},
			},
		))
		opts = append(opts, WithRequestEditorFn(BuffHeaders()))
	}
	return NewClientWithResponses(BuffURL, opts...)
}
