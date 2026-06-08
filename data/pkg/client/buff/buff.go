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
		req.Header.Set("User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")

		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")

		req.Header.Set("Referer", "https://buff.163.com/")
		req.Header.Set("Origin", "https://buff.163.com")

		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")

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
