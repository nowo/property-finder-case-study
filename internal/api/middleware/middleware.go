package middleware

import (
	"fmt"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/pkg/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CanPassWithToken(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	splittedAuthorization := strings.Split(authorization, " ")
	if len(splittedAuthorization) != 2 {
		fmt.Println("unauthorized user")
		return c.Status(fiber.StatusUnauthorized).JSON(general.Response{
			Status:  false,
			Message: messages.UNAUTHORIZED_USER.Error(),
		})
	}
	token := splittedAuthorization[1]
	_jwt := jwt.New().SetToken(token).DecodeToken()
	c.Locals("userID", _jwt.GetUserId())
	return c.Next()
}

func CantPassWithToken(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if len(authorization) != 0 || authorization != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(general.Response{
			Status:  false,
			Message: messages.AUTHORIZED_USER.Error(),
		})
	}

	return c.Next()
}

func SetContentTypeJSON(c *fiber.Ctx) error {
	c.Set("Content-type", "application/json; charset=utf-8")
	return c.Next()
}
