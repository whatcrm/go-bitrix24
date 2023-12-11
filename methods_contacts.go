package b24

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-bitrix24/models"
)

func (c *Get) Contacts(contactID string) (out []models.ContactResult, err error) {
	c.b24.log("GetContacts request is started...")
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmContactGet,
		In: &RequestParams{
			ID: contactID,
		},
		Out:    &models.Contact{},
		Params: nil,
	}

	if contactID != "" {
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = []models.ContactResult{options.Out.(*models.Contact).Result}
	}

	if contactID == "" {
		options.In = nil
		options.BaseURL = CrmContactList
		options.Out = &models.ContactList{}
		if err = c.b24.callMethod(options); err != nil {
			return
		}
		out = options.Out.(*models.ContactList).Result
	}
	return
}

func (c *Update) Contacts(id string, in *models.UpdateFields) (out MainResult, err error) {
	c.b24.log("UpdateContacts request is started...")
	contact := models.CompanyResult{
		ID:     id,
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmContactUpdate,
		In:      contact,
		Out:     &out,
		Params:  nil,
	}

	err = c.b24.callMethod(options)
	return
}

func (c *Create) Contacts(in *models.UpdateFields) (resp UFResult, err error) {
	c.b24.log("CreateContacts request is started...")
	contact := models.CompanyResult{
		Fields: in,
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: CrmContactAdd,
		In:      contact,
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

//func (c *Create) Contact(in []models.Contact) (out models.Contact, err error) {
//	c.b24.log("CreateContact request is started...")
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPost,
//		BaseURL: CrmContactAdd,
//		In:      in,
//		Out:     &out,
//		Params:  nil,
//	}
//
//	if err = c.b24.callMethod(options); err != nil {
//		return
//	}
//
//	c.b24.log("returning the struct...")
//	return
//}
//
//func (c *Update) Contacts(contactID string, in []models.Contact) (out []models.Contact, err error) {
//	c.b24.log("ModifyContacts request is started...")
//
//	options := callMethodOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: CrmContactUpdate,
//		In:      in,
//		Out:     &models.Contact{},
//		Params:  nil,
//	}
//
//	if contactID != "" {
//		options.BaseURL += "/" + contactID
//		options.In = in[0]
//		if err = c.b24.callMethod(options); err != nil {
//			return
//		}
//		out = []models.Contact{*options.Out.(*models.Contact)}
//	}
//
//	if contactID == "" {
//		options.Out = &models.Contact{}
//		if err = c.b24.callMethod(options); err != nil {
//			return
//		}
//		out = options.Out.(*models.RequestResponse).Embedded.Contacts
//	}
//
//	c.b24.log("returning the struct...")
//	return
//}
