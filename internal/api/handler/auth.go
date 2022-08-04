package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	domain "property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity"
	"property-finder-go-bootcamp-homework/internal/domain/user/service"
	"property-finder-go-bootcamp-homework/pkg/validation"
	"property-finder-go-bootcamp-homework/pkg/validation/user_info_validation"
)

func Register(c *fiber.Ctx) error {
	var userInfo entity.UserInfo

	encodeError := json.Unmarshal(c.Body(), &userInfo)
	if encodeError != nil {
		fmt.Println(encodeError)
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}

	validationError := user_info_validation.Validate(&userInfo)
	if validationError != nil {
		fmt.Println(validationError)
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    validationError,
		})
	}

	userService := service.New()
	token, registerError := userService.Register(domain.User{UserInfo: userInfo})
	if registerError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(general.Response{
			Status:  false,
			Message: registerError.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.REGISTER_SUCCESS,
		Data:    token,
	})
}

func Login(c *fiber.Ctx) error {
	var userInfo auth.LoginRequest
	encodeError := json.Unmarshal(c.Body(), &userInfo)
	if encodeError != nil {
		fmt.Println(encodeError)
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}

	validationError := validation.ValidateLoginRequest(&userInfo)
	if validationError != nil {
		fmt.Println(validationError)
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    validationError.Error(),
		})
	}

	userService := service.New()
	token, registerError := userService.Login(userInfo)
	if registerError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(general.Response{
			Status:  false,
			Message: registerError.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.LOGIN_SUCCESS.Error(),
		Data:    token,
	})
}
