package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
	"property-finder-go-bootcamp-homework/pkg/str"
	"regexp"
)

var Email = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(1, 50).Error(messages.WRONG_LENGHT.Error()),
	is.Email.Error(messages.WRONG_FORMAT.Error()),
}

var Firstname = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(1, 25).Error(messages.WRONG_LENGHT.Error()),
	validation.Match(regexp.MustCompile(str.NameRegex)).Error(messages.WRONG_FORMAT.Error()),
}

var Lastname = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(1, 25).Error(messages.WRONG_LENGHT.Error()),
	validation.Match(regexp.MustCompile(str.NameRegex)).Error(messages.WRONG_FORMAT.Error()),
}

var Password = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(8, 16).Error(messages.WRONG_LENGHT.Error()),
}

func Validate(dto interface{}) error {
	str := dto.(*entity_user.UserInfo)

	return validation.ValidateStruct(str,
		validation.Field(&str.Firstname, Firstname...),
		validation.Field(&str.Lastname, Lastname...),
		validation.Field(&str.Email, Email...),
		validation.Field(&str.Password, Password...),
	)
}
