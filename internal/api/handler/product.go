package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/product/service"
	"strconv"
)

func ListProducts(c *fiber.Ctx) error {
	productService := service.New()
	products, err := productService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: "",
		Data:    products,
	})

}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Query("id")
	productService := service.New()
	fmt.Println(id)
	integerId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("str conv error")
		return c.Status(fiber.StatusInternalServerError).JSON(general.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	product, err := productService.GetByID(integerId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(general.Response{
			Status:  false,
			Message: messages.PRODUCT_NOT_FOUND.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: "",
		Data:    product,
	})
}
