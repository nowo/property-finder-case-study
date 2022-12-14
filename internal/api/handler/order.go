package handler

import (
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/service_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
	"property-finder-go-bootcamp-homework/internal/domain/order/service_order"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/pkg/logger"
)

//Create order api handler. Complete order from carts of user. It also calculate total price of order.
func CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	cartService := service_cart.New(&repository_cart.CartRepository{}, &repository_product.ProductRepository{}, &repository_order.OrderRepository{})
	orderService := service_order.New(&repository_cart.CartRepository{}, &repository_order.OrderRepository{})
	productList, err := cartService.GetCartByUserID(uint(userID))
	if len(productList) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.CART_EMPTY.Error(),
			Data:    nil,
		})
	}
	if err != nil {
		logger.Errorf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	totalPrice, vatOfCart := cartService.CalculatePrice(productList, uint(userID))

	err = orderService.CreateOrder(uint(userID), totalPrice, vatOfCart)
	if err != nil {
		logger.Errorf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.ORDER_CREATE_SUCCESS,
		Data:    nil,
	})
}

//List user's all orders
func ListOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	orderService := service_order.New(&repository_cart.CartRepository{}, &repository_order.OrderRepository{})
	orderList, err := orderService.GetOrderByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.ORDER_LIST_SUCCESS,
		Data:    orderList,
	})
}
