package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/jwt/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/service_cart"
	"strconv"
)

func AddToCart(c *fiber.Ctx) error {
	productIDString := c.Query("id")

	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		fmt.Println("erdal")
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	userID := c.Locals("userID").(float64)
	addToCartError := service_cart.New().AddToCart(uint(userID), uint(productID))
	if addToCartError != nil {
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

func DeleteFromCart(c *fiber.Ctx) error {
	productIDString := c.Query("id")

	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		fmt.Println("erdal")
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	userID := c.Locals("userID").(float64)
	deleteFromCartError := service_cart.New().DeleteFromCart(uint(userID), uint(productID))
	if deleteFromCartError != nil {
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

func ListCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)
	fmt.Println("userID", userID)
	cartList, listCartError := service_cart.New().GetCartByUserID(uint(userID))
	if listCartError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: listCartError.Error(),
			Data:    nil,
		})
	}

	totalPrice, vatOfCart, err := service_cart.New().CalculatePrice(cartList)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.PRODUCT_LIST_CART_SUCCESS,
		Data: cart.Basket{
			Cart:       cartList,
			TotalPrice: totalPrice,
			VatOfCart:  vatOfCart,
		},
	})
}
