package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Leads(id string, params *RequestParams) (lead []models.LeadResult, err error) {
	c.b24.log("GetLead request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: CrmLeadGet,
		In:      nil,
		Out:     &models.Lead{},
		Params:  params,
	}

	if id != "" {
		// ID exists
		options.BaseURL += "/" + id
		err = c.b24.callMethod(options)
		if err != nil {
			return
		}
		lead = []models.LeadResult{options.Out.(*models.Lead).Result}
	}

	if id == "" {
		// All leads
		options.Out = &models.LeadList{}
		err = c.b24.callMethod(options)
		if err != nil {
			return
		}
		lead = options.Out.(*models.LeadList).Result
	}

	c.b24.log("returning the struct...")
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
