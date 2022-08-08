package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"property-finder-go-bootcamp-homework/dto/auth"
)

func ValidateLoginRequest(dto interface{}) error {
	str := dto.(*auth.LoginRequest)
	return validation.ValidateStruct(str,
		validation.Field(&str.Email, Email...),
		validation.Field(&str.Password, Password...),
	)
}
