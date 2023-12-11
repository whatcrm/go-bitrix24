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

func (c *Update) Companies(in *models.CompanyResult) (out MainResult, err error) {
	c.b24.log("UpdateCompanies request is started...")
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmCompanyUpdate,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) Companies(in []models.CompanyResult) (out []models.Company, err error) {
	c.b24.log("CustomersMode request is started...")

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmCompanyAdd,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	if err = c.b24.callMethod(options); err != nil {
		return
	}
	out = []models.Company{*options.Out.(*models.Company)}
	c.b24.log("returning the struct...")
	return
}
