package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Invoices(id string) (out models.InvoiceResult, err error) {
	// https://b24-4fyt5v.bitrix24.ru/rest/crm.item.get?auth=2691d36400645fee0065d9be00000001302a07baadfd547f28abce0a1223c63a76054b&entityTypeId=31&id=2
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

	err = c.b24.callMethod(options)
	out = options.Out.(*models.Invoice).Result
	return
}
