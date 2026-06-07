package buff

import (
	"fmt"
	"net/http"
	"net/url"
)

func WithProxy(proxy string) (ClientOption, error) {
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return nil, fmt.Errorf("%w proxy url %s", err, proxy)
	}

	return WithHTTPClient(
		&http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				Proxy:               http.ProxyURL(proxyUrl),
			},
		},
	), nil
}
