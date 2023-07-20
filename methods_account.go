package b24

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Create) TokenUPD(refreshToken string) (t *UpdatedTokens, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: "",
		In:      nil,
		Out:     &t,
		Params: &RequestParams{
			RefreshToken: refreshToken,
		},
	}

	err = c.b24.callMethod(options)
	return
}

func (b24 *API) IsAdmin() (out MainResult, err error) {

	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: UserAdmin,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	err = b24.callMethod(options)
	return
}

func (b24 *API) PlacementUnBind(handler, placement string) (out UnBind, err error) {
	// handler not required, if empty - all handlers will be removed
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: PlacementUnBind,
		In: &RequestParams{
			Placement: placement,
			Handler:   handler,
			// TODO LANG_ALL
			// https://dev.1c-bitrix.ru/rest_help/application_embedding/metods/placement_bind.php
		},
		Out:    &out,
		Params: nil,
	}

	err = b24.callMethod(options)
	return
}

func (b24 *API) PlacementBind(title, handler, placement string) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: PlacementBind,
		In: &RequestParams{
			Placement: placement,
			Handler:   handler,
			Title:     title,
			// TODO LANG_ALL
			// https://dev.1c-bitrix.ru/rest_help/application_embedding/metods/placement_bind.php
		},
		Out:    &out,
		Params: nil,
	}

	err = b24.callMethod(options)
	return
}

func (b24 *API) CallBind(event, handler string) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: EventBind,
		In: &RequestParams{
			Event:    event,
			Handler:  handler,
			AuthType: 0,
		},
		Out:    &out,
		Params: nil,
	}

	err = b24.callMethod(options)
	return
}

func (c *Create) UserFieldTypeAdd(userTypeID, handler, title, description string) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldTypeAdd,
		In: &RequestParams{
			UserTypeID:  userTypeID,
			Handler:     handler,
			Title:       title,
			Description: description,
		},
		Out:    &out,
		Params: nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) UserFieldAdd(base, userTypeID, fieldName, efl, description string) (out MainResult, err error) {
	if base != CrmContactUserField && base != CrmLeadUserField &&
		base != CrmDealUserField && base != CrmCompanyUserField {
		return out, fiber.NewError(fiber.StatusForbidden, "wrong base url")
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: base,
		In: &RequestParams{
			UserTypeID:    userTypeID,
			FieldName:     fieldName,
			EditFormLabel: efl,
			Description:   description,
		},
		Out:    &out,
		Params: nil,
	}

	err = c.b24.callMethod(options)
	return
}
