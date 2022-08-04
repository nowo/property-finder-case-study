package user_info_validation

import (
	"property-finder-go-bootcamp-homework/src/config/messages"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var Email = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(1, 50).Error(messages.WRONG_LENGHT.Error()),
	is.Email.Error(messages.WRONG_FORMAT.Error()),
}
