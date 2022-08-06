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
	api := app.Group(endpoints.API_VERSION, middleware.SetContentTypeJSON)
	auth := api.Group(endpoints.AUTH_ENDPOINT)
	cart := api.Group(endpoints.CART_ENDPOINT)
	//product := api.Group(endpoints.PRODUCT_ENDPOINT)
	order := api.Group(endpoints.ORDER_ENDPOINT)
	auth.Post(endpoints.REGISTER_ENDPOINT, middleware.TokenCantGo, handler.Register)
	auth.Post(endpoints.LOGIN_ENDPOINT, middleware.TokenCantGo, handler.Login)

	api.Get(endpoints.PRODUCTS_ENDPOINT, handler.ListProducts)
	api.Get(endpoints.PRODUCT_ENDPOINT, handler.GetProductByID)

	cart.Post(endpoints.EMPTY, middleware.TokenCanGo, handler.AddToCart)
	cart.Post(endpoints.DELETE_ENDPOINT, middleware.TokenCanGo, handler.DeleteFromCart)
	cart.Get(endpoints.EMPTY, middleware.TokenCanGo, handler.ListCart)

	order.Post(endpoints.CREATE_ENDPOINT, middleware.TokenCanGo, handler.CreateOrder)
	order.Get(endpoints.LIST_ENDPOINT, middleware.TokenCanGo, handler.ListOrders)
	log.Println(app.Listen(endpoints.PORT))
}
