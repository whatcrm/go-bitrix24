package b24

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		client:       &http.Client{},
	}
}

func NewWebhook(host string) (*API, error) {
	if !strings.HasPrefix(host, "http") {
		return nil, fmt.Errorf("invalid webhook URL")
	}

	if _, err := url.Parse(host); err != nil {
		return nil, fmt.Errorf("invalid webhook URL: %w", err)
	}

	return &API{
		WebhookURL: host,
		client:     &http.Client{},
	}, nil
}

func (b24 *API) SetOptions(domain, auth string, debug bool) error {
	b24.Domain = domain

	if auth == "" {
		return fmt.Errorf("accessToken is not set")
	}

	b24.Auth = auth
	b24.Debug = debug
	return nil
}

func (b24 *API) SetFallback(fallbackRefresh string) {
	b24.FallbackRefreshToken = fallbackRefresh
}

func (b24 *API) SetProxy(proxy string) error {
	if proxy == "" {
		b24.client.Transport = http.DefaultTransport.(*http.Transport).Clone()
		return nil
	}

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.Proxy = http.ProxyURL(proxyURL)

	b24.client.Transport = tr
	return nil
}

func (b24 *API) Close() {
	if tr, ok := b24.client.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
	b24.client = nil
}
