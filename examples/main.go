package main

import (
	"github.com/gofiber/fiber/v2"
	goBX24 "github.com/whatcrm/go-bitrix24"

	"log"
)

const (
	contactID = "5"

	clientID     = "app.clientID"
	clientSecret = "clientSecret"
	domain       = "portal.bitrix24.ru"
	auth         = "auth"
)

func main() {
	server()
}

func server() {
	app := fiber.New()
	app.Get("", handler)
	log.Fatal(app.Listen(":3000"))
}

func handler(ctx *fiber.Ctx) error {
	log.Println(domain)
	b24 := goBX24.NewAPI(clientID, clientSecret)

	if err := b24.SetOptions(domain, auth, true); err != nil {
		return err
	}

	admin, _ := b24.IsAdmin()
	log.Println(admin.Result)

	res, err := b24.Get().Contacts(contactID)
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}
