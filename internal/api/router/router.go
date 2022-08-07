package router

import (
	"log"
	"property-finder-go-bootcamp-homework/internal/.config"
	"property-finder-go-bootcamp-homework/internal/api/handler"
	"property-finder-go-bootcamp-homework/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()
	api := app.Group(_config.API_VERSION, middleware.SetContentTypeJSON)
	auth := api.Group(_config.AUTH_ENDPOINT)
	cart := api.Group(_config.CART_ENDPOINT)
	product := api.Group(_config.PRODUCT_ENDPOINT)
	order := api.Group(_config.ORDER_ENDPOINT)

	auth.Post(_config.REGISTER_ENDPOINT, middleware.CantPassWithToken, handler.RegisterUser)
	auth.Post(_config.LOGIN_ENDPOINT, middleware.CantPassWithToken, handler.Login)

	product.Get(_config.LIST_ENDPOINT, handler.ListProducts)
	product.Get(_config.EMPTY, handler.GetProductByID)

	cart.Post(_config.EMPTY, middleware.CanPassWithToken, handler.AddToCart)
	cart.Post(_config.DELETE_ENDPOINT, middleware.CanPassWithToken, handler.DeleteFromCart)
	cart.Get(_config.EMPTY, middleware.CanPassWithToken, handler.ListCart)

	order.Post(_config.CREATE_ENDPOINT, middleware.CanPassWithToken, handler.CreateOrder)
	order.Get(_config.LIST_ENDPOINT, middleware.CanPassWithToken, handler.ListOrders)

	log.Println(app.Listen(_config.PORT))
}
