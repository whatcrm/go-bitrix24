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
		return
	}

	options.In = nil
	options.BaseURL = CrmDealList
	options.Out = &models.DealList{}
	if err = c.b24.callMethod(options); err != nil {
		return
	}
	out = options.Out.(*models.DealList).Result

	c.b24.log("returning the struct...")
	return
}

func (c *Update) Deals(id string, in *models.UpdateFields) (MainResult, error) {
	out := MainResult{}
	deal := models.DealResult{
		ID:     id,
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDealUpdate,
		In:      deal,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Create) Deals(in *models.UpdateFields) (UFResult, error) {
	c.b24.log("CreateDeals request is started...")
	resp := UFResult{}
	deal := models.DealResult{
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmDealAdd,
		In:      deal,
		Out:     &resp,
		Params:  nil,
	}

	return resp, c.b24.callMethod(options)
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
