package user_info_validation

import (
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/pkg/str"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var Firstname = []validation.Rule{
	validation.Required.Error(messages.REQUIRED_FIELD.Error()),
	validation.Length(1, 25).Error(messages.WRONG_LENGHT.Error()),
	validation.Match(regexp.MustCompile(str.NameRegex)).Error(messages.WRONG_FORMAT.Error()),
}
