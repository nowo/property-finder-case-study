package user_info_validation

import (
	"property-finder-go-bootcamp-homework/internal/domain/user/entity"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(dto interface{}) error {
	str := dto.(*entity.UserInfo)

	return validation.ValidateStruct(str,
		validation.Field(&str.Firstname, Firstname...),
		validation.Field(&str.Lastname, Lastname...),
		validation.Field(&str.Email, Email...),
		validation.Field(&str.Password, Password...),
	)
}
