package auth_handler

import (
	"encoding/json"
	"fmt"
	"property-finder-go-bootcamp-homework/src/config/messages"
	domain "property-finder-go-bootcamp-homework/src/domain/user"
	"property-finder-go-bootcamp-homework/src/domain/user/entity"
	"property-finder-go-bootcamp-homework/src/domain/user/service"
	"property-finder-go-bootcamp-homework/src/dto/general"
	"property-finder-go-bootcamp-homework/src/pkg/validation/user_info_validation"

	"github.com/gofiber/fiber/v2"
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
			Data:    validationError.Error(),
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
