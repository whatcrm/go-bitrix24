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
	domain       = "b24-jh8et8.bitrix24.ru"
	auth         = "0119b96400645fee0065854e00000001302a0751f64bf061a7f786040aa08917c5bd4b"
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

	contacts, _ := b24.Get().Contacts(contactID)
	return ctx.JSON(contacts)
}
