package b24

func (b24 *API) fallback() bool {
	if b24.FallbackRefreshToken == "" {
		return false
	}
	upd, err := b24.Create().TokenUPD(b24.FallbackRefreshToken)
	if err != nil {
		return false
	}

	if err = b24.SetOptions(b24.Domain, upd.AccessToken, b24.Debug); err != nil {
		return false
	}

	return true
}
