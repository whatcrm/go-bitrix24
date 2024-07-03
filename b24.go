package b24

import (
	"fmt"
)

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
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
