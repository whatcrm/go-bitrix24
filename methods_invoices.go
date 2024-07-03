package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Invoices(id string) (models.InvoiceResult, error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmItemGet,
		In: &RequestParams{
			EntityTypeID: "31",
			Id:           id,
		},
		Out:    &models.Invoice{},
		Params: nil,
	}

	return options.Out.(*models.Invoice).Result, c.b24.callMethod(options)
}
