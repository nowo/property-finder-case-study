package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/jwt/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/service_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/pkg/logger"
	"strconv"
)

//Add product to cart api handler. Takes id of product as query parameter. Returns status code 200 if success.
func AddToCart(c *fiber.Ctx) error {
	productIDString := c.Query("id")
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		logger.Errorf(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	userID := c.Locals("userID").(float64)
	addToCartError := service_cart.New(&repository_cart.CartRepository{}, &repository_product.ProductRepository{}, &repository_order.OrderRepository{}).AddToCart(uint(userID), uint(productID))
	if addToCartError != nil {
		logger.Errorf(addToCartError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: addToCartError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.PRODUCT_ADD_TO_CART_SUCCESS,
		Data:    nil,
	})
}

//Delete product from cart api handler. Takes id of product as query parameter. Returns status code 200 if success.
func DeleteFromCart(c *fiber.Ctx) error {
	productIDString := c.Query("id")

	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		logger.Errorf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	userID := c.Locals("userID").(float64)
	deleteFromCartError := service_cart.New(&repository_cart.CartRepository{}, &repository_product.ProductRepository{}, &repository_order.OrderRepository{}).DeleteFromCart(uint(userID), uint(productID))
	if deleteFromCartError != nil {
		logger.Errorf(deleteFromCartError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: deleteFromCartError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.PRODUCT_DELETE_FROM_CART_SUCCESS,
		Data:    nil,
	})
}

//Get cart api handler. Returns cart list and prices of cart if it success
func ListCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	cartList, listCartError := service_cart.New(&repository_cart.CartRepository{}, &repository_product.ProductRepository{}, &repository_order.OrderRepository{}).GetCartByUserID(uint(userID))
	if listCartError != nil {
		logger.Errorf(listCartError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: listCartError.Error(),
			Data:    nil,
		})
	}
	totalPrice, vatOfCart := service_cart.New(&repository_cart.CartRepository{}, &repository_product.ProductRepository{}, &repository_order.OrderRepository{}).CalculatePrice(cartList, uint(userID))

	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.PRODUCT_LIST_CART_SUCCESS,
		Data: general.BasketResponse{
			Cart:       cartList,
			TotalPrice: totalPrice,
			VatOfCart:  vatOfCart,
		},
	})
}
