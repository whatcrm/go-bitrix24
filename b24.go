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

func (b24 *API) SetOptions(domain, auth string, debug bool) error {
	b24.Domain = domain

	// Для вебхуков проверяем формат
	if strings.HasPrefix(auth, "http") {
		if _, err := url.Parse(auth); err != nil {
			return fmt.Errorf("invalid webhook URL: %w", err)
		}
	} else if auth == "" {
		return fmt.Errorf("accessToken is not set")
	}

	b24.Auth = auth
	b24.Debug = debug
	return nil
}

func (b24 *API) SetFallback(fallbackRefresh string) {
	b24.FallbackRefreshToken = fallbackRefresh
}
