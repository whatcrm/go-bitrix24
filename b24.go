package b24

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var sharedTransport = &http.Transport{
	TLSHandshakeTimeout: 10 * time.Second,
	MaxIdleConns:        100,
	IdleConnTimeout:     90 * time.Second,
}

var sharedClient = &http.Client{
	Transport: sharedTransport,
	Timeout:   60 * time.Second,
}

type proxyKey struct{}

func init() {
	sharedTransport.Proxy = func(req *http.Request) (*url.URL, error) {
		if v := req.Context().Value(proxyKey{}); v != nil {
			if u, ok := v.(*url.URL); ok {
				return u, nil
			}
		}
		return nil, nil // no proxy
	}
}

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		client:       sharedClient,
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

func (b24 *API) SetProxy(proxy string) {
	b24.Proxy = proxy
	return
}

func (b24 *API) SetTimeout(d time.Duration) {
	sharedClient.Timeout = d
}

func (b24 *API) Close() {
	if tr, ok := b24.client.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
}
