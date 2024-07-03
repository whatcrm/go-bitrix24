package b24

func (b24 *API) fallback(refresh string) error {
	upd, err := b24.Create().TokenUPD(refresh)
	if err != nil {
		return err
	}

	if err = b24.SetOptions(upd.Domain, upd.AccessToken, b24.Debug); err != nil {
		return err
	}

	return nil
}
