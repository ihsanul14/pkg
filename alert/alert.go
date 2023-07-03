package alert

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	HTTP_POST       = "POST"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

type Alert interface {
	Send() error
	BuildContent() string
}

func GenerateClient(proxies []string) *http.Client {
	if proxies == nil {
		return &http.Client{}
	}

	proxyURLs := proxies
	proxyFunc := func(req *http.Request) (*url.URL, error) {
		if len(proxyURLs) == 0 {
			return nil, fmt.Errorf("no proxies available")
		}

		proxyURL := proxyURLs[0]
		proxyURLs = proxyURLs[1:]

		url, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		return url, nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxyFunc,
		},
	}

	return client
}
