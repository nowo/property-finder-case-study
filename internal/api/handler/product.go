package handler

import (
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/internal/domain/product/service_product"
	"property-finder-go-bootcamp-homework/pkg/logger"
	"strconv"
)

//List all products api handler. Returns products list.
func ListProducts(c *fiber.Ctx) error {
	productService := service_product.New(&repository_product.ProductRepository{})
	products, err := productService.GetAll()
	if err != nil {
		logger.Errorf(err.Error())
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

//Get product by id api handler. Takes id of product as query parameter. Returns product.
func GetProductByID(c *fiber.Ctx) error {
	id := c.Query("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}
	productService := service_product.New(&repository_product.ProductRepository{})

	product, getByIDError := productService.GetByID(uint(productID))
	if getByIDError != nil {
		logger.Errorf(getByIDError.Error())
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
