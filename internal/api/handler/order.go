package handler

import (
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart/service_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/service_order"
	"property-finder-go-bootcamp-homework/pkg/logger"
)

func CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)

	cartService := service_cart.New()
	orderService := service_order.New()
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

func ListOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	orderService := service_order.New()
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
