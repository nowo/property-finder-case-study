package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/pkg/validation/user_info_validation"
)

func ValidateLoginRequest(dto interface{}) error {
	str := dto.(*auth.LoginRequest)

	return validation.ValidateStruct(str,
		validation.Field(&str.Email, user_info_validation.Email...),
		validation.Field(&str.Password, user_info_validation.Password...),
	)
}
