package b24

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
	"log"
)

func (c *Create) TokenUPD(refreshToken string) (t UpdatedTokens, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodGet,
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

func (b24 *API) CallUnBind(event, handler string) (out UnBind, err error) {
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

	err = b24.callMethod(options)
	return
}

func (c *Create) UserFieldType(in *models.UserField) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldTypeAdd,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Delete) UserFieldType(in *models.UserField) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldTypeDelete,
		In:      &in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) UserField(in *models.UserField, baseURL string) (out UFResult, err error) {
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

	err = c.b24.callMethod(options)
	return
}

func (c *Get) UserField(baseURL string) (out UserfieldList, err error) {
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

	err = c.b24.callMethod(options)
	return
}

func (c *Delete) UserField(id string, baseURL string) (out MainResult, err error) {
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

	err = c.b24.callMethod(options)
	return
}

func (c *Get) UserFieldConfig(in *RequestParams) (out UserFieldConfig, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldConfigList,
		In:      nil,
		Out:     &out,
		Params:  in,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Delete) UserFieldConfig(in *RequestParams) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: UserFieldConfigDelete,
		In:      nil,
		Out:     &out,
		Params:  in,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Get) FindDuplicates(in *DuplicatesParams) (out DuplicatesResponse, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDuplicatesFindByComm,
		In:      in,
		Out:     &DuplicatesNotFound{},
		Params:  nil,
	}

	err = c.b24.callMethod(options)

	test := options.Out.(*DuplicatesNotFound).Result

	out, err = caseDuplicatesNotFound(test)
	if err != nil {
		c.b24.log("duplicates are found...")
		out, err = caseDuplicatesResponse(test)
	}

	return
}

func caseDuplicatesNotFound(test any) (DuplicatesResponse, error) {

	out := DuplicatesResponse{}

	val, ok := test.(interface{})
	if !ok {
		return out, fmt.Errorf("any is not an interface")
	}
	m, err := json.Marshal(val)

	if err = json.Unmarshal(m, &out.CONTACT); err != nil {
		return out, err
	}
	return out, nil

}

func caseDuplicatesResponse(test any) (DuplicatesResponse, error) {
	out := DuplicatesResponse{}

	val, ok := test.(map[string]interface{})
	if !ok {
		fmt.Println("any is not a map")
		return out, nil
	}

	m, err := json.Marshal(val)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(m, &out); err != nil {
		log.Println(err)
		return out, err
	}
	return out, nil
}
