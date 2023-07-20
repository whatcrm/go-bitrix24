package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Companies(companyID string, params *RequestParams) (out []models.CompanyResult, err error) {
	c.b24.log("CustomersMode request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: CrmCompanyGet,
		In:      nil,
		Out:     &models.Company{},
		Params:  params,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		if err = c.b24.callMethod(options); err != nil {
			return
		}

		out = []models.CompanyResult{options.Out.(*models.Company).Result}
	}

	if companyID == "" {
		options.BaseURL = CrmCompanyList
		options.Out = &models.CompanyList{}
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = options.Out.(*models.CompanyList).Result
	}

	c.b24.log("returning the struct...")
	return
}

//
//func (c *Create) Company(in *[]models.Company) (out []models.Company, err error) {
//	c.b24.log("CustomersMode request is started...")
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPost,
//		BaseURL: CrmCompanyCreate,
//		In:      in,
//		Out:     &models.Company{},
//		Params:  nil,
//	}
//
//	if len(*in) <= 1 {
//		options.Out = &models.Company{}
//		if err = c.b24.makeRequest(options); err != nil {
//			return
//		}
//		out = []models.Company{*options.Out.(*models.Company)}
//	} else {
//		options.Out = &models.RequestResponse{}
//		if err = c.b24.makeRequest(options); err != nil {
//			return
//		}
//		out = options.Out.(*models.RequestResponse).Embedded.Companies
//	}
//
//	c.b24.log("returning the struct...")
//	return
//}
//
//func (c *Update) Company(companyID string, in *[]models.Company) (out []models.Company, err error) {
//	c.b24.log("CustomersMode request is started...")
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: companiesURL,
//		In:      in,
//		Out:     &models.Company{},
//		Params:  nil,
//	}
//
//	if companyID != "" {
//		options.BaseURL += "/" + companyID
//		options.In = (*in)[0]
//
//		if err = c.b24.makeRequest(options); err != nil {
//			return
//		}
//		out = []models.Company{*options.Out.(*models.Company)}
//	}
//
//	if companyID == "" {
//		options.Out = &models.RequestResponse{}
//		if err = c.b24.makeRequest(options); err != nil {
//			return
//		}
//		out = options.Out.(*models.RequestResponse).Embedded.Companies
//	}
//
//	c.b24.log("returning the struct...")
//	return
//}
