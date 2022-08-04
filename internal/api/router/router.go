package router

import (
	"log"
	"property-finder-go-bootcamp-homework/internal/.config/endpoints"
	"property-finder-go-bootcamp-homework/internal/api/handler"
	"property-finder-go-bootcamp-homework/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()
	auth := app.Group(endpoints.AUTH_ENDPOINT)
	auth.Post(endpoints.REGISTER_ENDPOINT, middleware.TokenCantGo, handler.Register)
	auth.Post(endpoints.LOGIN_ENDPOINT, middleware.TokenCantGo, handler.Login)
	app.Get(endpoints.GET_PRODUCTS_ENDPOINT, handler.ListProducts)
	app.Get(endpoints.GET_PRODUCT, handler.GetProductByID)
	log.Fatal(app.Listen(endpoints.PORT))
}
