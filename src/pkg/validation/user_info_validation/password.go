package user_info_validation

import (
	"property-finder-go-bootcamp-homework/src/config/messages"

	validation "github.com/go-ozzo/ozzo-validation"
)

var Password = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(8, 16).Error(messages.WRONG_LENGHT.Error()),
}
