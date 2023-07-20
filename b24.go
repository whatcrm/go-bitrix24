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
	if isRegex(domain) {
		b24.Domain = domain
	} else {
		return fmt.Errorf("domain name is not set")
	}

	if auth != "" {
		b24.Auth = auth
	} else {
		return fmt.Errorf("accessToken is not set")
	}

	b24.Debug = debug

	return nil
}
