package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Deals(id string) (out []models.DealResult, err error) {
	c.b24.log("GetDeals request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDealGet,
		In: &RequestParams{
			ID: id,
		},
		Out:    &models.Deal{},
		Params: nil,
	}

	if id != "" {
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = []models.DealResult{options.Out.(*models.Deal).Result}
	}

	if id == "" {
		options.In = nil
		options.BaseURL = CrmDealList
		options.Out = &models.DealList{}
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = options.Out.(*models.DealList).Result
	}

	c.b24.log("returning the struct...")
	return
}

func (c *Update) Deals(in models.DealResult) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDealUpdate,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) Deals(params *models.DealResult) (resp UFResult, err error) {
	c.b24.log("CreateDeals request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDealAdd,
		In:      params,
		Out:     &resp,
		Params:  nil,
	}
	err = c.b24.callMethod(options)
	if err != nil {
		return
	}

	c.b24.log("returning the struct...")
	return
}

//
//func (c *Update) Lead(id string, lead *models.Lead, params *RequestParams) (resp []models.LeadResult, err error) {
//	c.b24.log("ModifyLead request started...")
//
//	params.ID = id
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: CrmLeadUpdate,
//		In:      lead,
//		Out:     &resp,
//		Params:  params,
//	}
//	err = c.b24.callMethod(options)
//	if err != nil {
//		return
//	}
//
//	c.b24.log("returning the struct...")
//	return
//}
