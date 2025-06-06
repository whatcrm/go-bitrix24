package b24

import (
	"fmt"
	"net/url"
	"strings"
)

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
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
