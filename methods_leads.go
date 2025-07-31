package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Leads(id string) (lead []models.LeadResult, err error) {
	c.b24.log("GetLead request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmLeadGet,
		In: &RequestParams{
			ID: id,
		},
		Out:    &models.Lead{},
		Params: nil,
	}

	if id != "" {
		// ID exists
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		lead = []models.LeadResult{options.Out.(*models.Lead).Result}
		return
	}

	options.In = nil
	options.BaseURL = CrmLeadList
	options.Out = &models.LeadList{}
	if err = c.b24.callMethod(options); err != nil {
		return
	}
	lead = options.Out.(*models.LeadList).Result

	c.b24.log("returning the struct...")
	return
}

func (c *Get) LeadsWithParams(params *RequestParams) (out []models.LeadResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmLeadList,
		Out:     &models.LeadList{},
		Params:  params,
	}
	if err = c.b24.callMethod(options); err != nil {
		return
	}
	out = options.Out.(*models.LeadList).Result
	return
}

func (c *Update) Leads(id string, in *models.UpdateFields) (MainResult, error) {
	out := MainResult{}
	deal := models.DealResult{
		ID:     id,
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmLeadUpdate,
		In:      deal,
		Out:     &out,
		Params:  nil,
	}

	return out, c.b24.callMethod(options)
}

func (c *Create) Leads(in *models.UpdateFields) (UFResult, error) {
	resp := UFResult{}
	c.b24.log("CreateLeads request is started...")
	deal := models.DealResult{
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmLeadAdd,
		In:      deal,
		Out:     &resp,
		Params:  nil,
	}

	return resp, c.b24.callMethod(options)
}

//
//func (c *Create) Leads(lead *[]models.Lead, params *RequestParams) (resp []models.LeadResult, err error) {
//	c.b24.log("CreateLeads request is started...")
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPost,
//		BaseURL: CrmLeadAdd,
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
//	log.Println(lead)
//	return
//}
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
