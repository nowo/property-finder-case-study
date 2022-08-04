package router

import (
	"log"
	"property-finder-go-bootcamp-homework/src/config/endpoints"
	"property-finder-go-bootcamp-homework/src/fiber/handler/auth_handler"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()
	auth := app.Group(endpoints.AUTH_ENDPOINT)
	auth.Post(endpoints.REGISTER_ENDPOINT, auth_handler.Register)
	log.Fatal(app.Listen(endpoints.PORT))
}
