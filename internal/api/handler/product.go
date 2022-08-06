package handler

import (
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/product/service_product"
	"strconv"
)

func ListProducts(c *fiber.Ctx) error {
	productService := service_product.New()
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
		Message: messages.PRODUCT_LIST_SUCCESS,
		Data:    products,
	})

}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Query("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	productService := service_product.New()

	product, getByIDError := productService.GetByID(uint(productID))
	if getByIDError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: getByIDError.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.PRODUCT_LIST_SUCCESS,
		Data:    product,
	})
}
