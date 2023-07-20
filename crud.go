package b24

type Create struct {
	b24 *API
}
type Get struct {
	b24 *API
}
type Update struct {
	b24 *API
}
type Delete struct {
	b24 *API
}

func (b24 *API) Create() *Create {
	return &Create{b24}
}

func (b24 *API) Get() *Get {
	return &Get{b24}
}
func (b24 *API) Update() *Update {
	return &Update{b24}
}
func (b24 *API) Delete() *Delete {
	return &Delete{b24}
}
