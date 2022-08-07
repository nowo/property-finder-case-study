package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/internal/domain/user/service_user"
	"property-finder-go-bootcamp-homework/pkg/logger"
	"property-finder-go-bootcamp-homework/pkg/validation"
)

func RegisterUser(c *fiber.Ctx) error {
	var userInfo entity_user.UserInfo

	encodeError := json.Unmarshal(c.Body(), &userInfo)
	if encodeError != nil {
		logger.Errorf(encodeError.Error())

		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}

	validationError := validation.Validate(&userInfo)
	if validationError != nil {
		logger.Errorf(validationError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    validationError,
		})
	}

	userService := service_user.New()
	token, registerError := userService.
		Register(user.User{
			UserInfo: userInfo,
		})
	if registerError != nil {
		logger.Errorf(registerError.Error())
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
		logger.Errorf(encodeError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    nil,
		})
	}

	validationError := validation.ValidateLoginRequest(&userInfo)
	if validationError != nil {
		logger.Errorf(validationError.Error())
		return c.Status(fiber.StatusBadRequest).JSON(general.Response{
			Status:  false,
			Message: messages.BAD_REQUEST.Error(),
			Data:    validationError.Error(),
		})
	}

	userService := service_user.New()
	token, registerError := userService.Login(userInfo)
	if registerError != nil {
		logger.Errorf(registerError.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(general.Response{
			Status:  false,
			Message: registerError.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(general.Response{
		Status:  true,
		Message: messages.LOGIN_SUCCESS,
		Data:    token,
	})
}
