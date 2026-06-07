package buff

import (
	"context"
	"errors"
	"net/http"

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
		proxy, err := WithProxy(*cfg.Proxy)
		if err != nil {
			return nil, err
		}
		opts = append(opts, proxy)
	}
	opts = append(opts, WithRequestEditorFn(BuffHeaders()))
	return NewClientWithResponses(BuffURL, opts...)
}
