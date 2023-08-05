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
	}

	if id == "" {
		options.In = nil
		options.BaseURL = CrmLeadList
		options.Out = &models.LeadList{}
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		lead = options.Out.(*models.LeadList).Result
	}

	c.b24.log("returning the struct...")
	return
}

func (c *Update) Leads(in models.LeadResult) (out MainResult, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmLeadUpdate,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
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
