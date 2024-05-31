package b24

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Create) TokenUPD(refreshToken string) (UpdatedTokens, error) {
	t := UpdatedTokens{}
	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: "",
		In:      nil,
		Out:     &t,
		Params: &RequestParams{
			RefreshToken: refreshToken,
		},
	}

	return t, c.b24.callMethod(options)
}

func (b24 *API) IsAdmin() (MainResult, error) {
	out := MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: UserAdmin,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	return out, b24.callMethod(options)
}

func (b24 *API) PlacementUnBind(handler, placement string) (UnBind, error) {
	out := UnBind{}
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

	return out, b24.callMethod(options)
}

func (b24 *API) PlacementBind(title, handler, placement string) (MainResult, error) {
	out := MainResult{}
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

	return out, b24.callMethod(options)
}

func (b24 *API) CallBind(event, handler string) (MainResult, error) {
	out := MainResult{}
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

	return out, b24.callMethod(options)
}

func (b24 *API) CallUnBind(event, handler string) (UnBind, error) {
	out := UnBind{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: EventUnBind,
		In: &RequestParams{
			Event:   event,
			Handler: handler,
		},
		Out:    &out,
		Params: nil,
	}

	return out, b24.callMethod(options)
}

func (c *Create) UserFieldType(in *models.UserField) (MainResult, error) {
	out := MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldTypeAdd,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Delete) UserFieldType(in *models.UserField) (MainResult, error) {
	out := MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldTypeDelete,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Create) UserField(in *models.UserField, baseURL string) (UFResult, error) {
	out := UFResult{}
	if baseURL != CrmContactUserFieldAdd && baseURL != CrmLeadUserFieldAdd &&
		baseURL != CrmDealUserFieldAdd && baseURL != CrmCompanyUserFieldAdd {
		return out, fiber.NewError(fiber.StatusForbidden, "wrong base url: "+baseURL)
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: baseURL,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Get) UserField(baseURL string) (UserfieldList, error) {
	out := UserfieldList{}
	if baseURL != CrmContactUserFieldList && baseURL != CrmLeadUserFieldList &&
		baseURL != CrmDealUserFieldList && baseURL != CrmCompanyUserFieldList {
		return out, fiber.NewError(fiber.StatusForbidden, "wrong base url: "+baseURL)
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: baseURL,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Delete) UserField(id string, baseURL string) (MainResult, error) {
	out := MainResult{}
	if baseURL != CrmContactUserFieldDelete && baseURL != CrmLeadUserFieldDelete &&
		baseURL != CrmDealUserFieldDelete && baseURL != CrmCompanyUserFieldDelete {
		return out, fiber.NewError(fiber.StatusForbidden, "wrong base url: "+baseURL)
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: baseURL,
		In: &RequestParams{
			ID: id,
		},
		Out:    &out,
		Params: nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Get) UserFieldConfig(in *RequestParams) (UserFieldConfig, error) {
	out := UserFieldConfig{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldConfigList,
		In:      nil,
		Out:     &out,
		Params:  in,
	}

	return out, c.b24.callMethod(options)
}

func (c *Delete) UserFieldConfig(in *RequestParams) (MainResult, error) {
	out := MainResult{}
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldConfigDelete,
		In:      nil,
		Out:     &out,
		Params:  in,
	}

	return out, c.b24.callMethod(options)
}

func (c *Get) FindDuplicates(in *DuplicatesParams) (out DuplicatesResponse, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDuplicatesFindByComm,
		In:      in,
		Out:     &DuplicatesNotFound{},
		Params:  nil,
	}

	if err = c.b24.callMethod(options); err != nil {
		return
	}

	test, ok := options.Out.(*DuplicatesNotFound)
	if !ok {
		return out, errors.New("unable to assert to DuplicatesNotFound")
	}

	out, err = caseDuplicatesNotFound(test.Result)
	if err != nil {
		c.b24.log("duplicates are found...")
		out, err = caseDuplicatesResponse(test.Result)
	}

	return
}

func caseDuplicatesNotFound(test any) (DuplicatesResponse, error) {
	out := DuplicatesResponse{}
	m, err := json.Marshal(test)
	if err != nil {
		return out, err
	}

	return out, json.Unmarshal(m, &out.CONTACT)
}

func caseDuplicatesResponse(test any) (DuplicatesResponse, error) {
	out := DuplicatesResponse{}
	val, ok := test.(map[string]any)
	if !ok {
		return out, errors.New("unable to assert to DuplicatesResponse")
	}

	m, err := json.Marshal(val)
	if err != nil {
		return out, err
	}

	return out, json.Unmarshal(m, &out)
}
