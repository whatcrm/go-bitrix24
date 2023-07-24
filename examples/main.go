package main

import (
	"github.com/gofiber/fiber/v2"
	goBX24 "github.com/whatcrm/go-bitrix24"

	"log"
)

const (
	contactID = "5"

	clientID     = "app.647db175212ef1.23413044"
	clientSecret = "8DfQWjTSKWqB631Dk3ycyW3QawvAXeWKfhgfQfd8blJz6zdVqV"
	domain       = "b24-jh8et8.bitrix24.ru"
	auth         = "4c08be6400645fee0065854e0000000100000093e132249186df230b02e277ee8adec0"
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
