package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Companies(id string) (out []models.CompanyResult, err error) {
	c.b24.log("GetCompanies request is started...")
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmCompanyGet,
		In: &RequestParams{
			ID: id,
		},
		Out:    &models.Company{},
		Params: nil,
	}

	if id != "" {
		if err = c.b24.callMethod(options); err != nil {
			return
		}

		out = []models.CompanyResult{options.Out.(*models.Company).Result}
	}

	if id == "" {
		options.In = nil
		options.BaseURL = CrmCompanyList
		options.Out = &models.CompanyList{}
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = options.Out.(*models.CompanyList).Result
	}
	return
}

func (c *Update) Companies(id string, in *models.UpdateFields) (out MainResult, err error) {
	company := models.CompanyResult{
		ID:     id,
		Fields: in,
	}

	c.b24.log("UpdateCompanies request is started...")
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmCompanyUpdate,
		In:      company,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) Companies(in *models.UpdateFields) (resp UFResult, err error) {
	c.b24.log("CreateCompanies request is started...")
	company := models.CompanyResult{
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmCompanyAdd,
		In:      company,
		Out:     &resp,
		Params:  nil,
	}

	if err = c.b24.callMethod(options); err != nil {
		return
	}
	
	
	c.b24.log("returning the struct...")
	return
}