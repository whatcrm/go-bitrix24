package main

import (
	"github.com/gofiber/fiber/v2"
	goBX24 "github.com/whatcrm/go-bitrix24"

	"log"
)

const (
	contactID = "1"

	clientID     = "app.647db175212ef1.23413044"
	clientSecret = "8DfQWjTSKWqB631Dk3ycyW3QawvAXeWKfhgfQfd8blJz6zdVqV"
	domain       = "b24-0kwfc0.bitrix24.ru"
	auth         = "cfd3b86400645fee0065234c00000001302a0708c997494713f8f3096073da75a4a9d7"
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
	b24 := goBX24.NewAPI(clientID, clientSecret)

	if err := b24.SetOptions(domain, auth, false); err != nil {
		return err
	}

	admin, _ := b24.IsAdmin()
	log.Println(admin.Result)

	contacts, _ := b24.Get().Contacts(contactID)
	return ctx.JSON(contacts)
}
