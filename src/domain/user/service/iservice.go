package service

import (
	domain "property-finder-go-bootcamp-homework/src/domain/user"
	"property-finder-go-bootcamp-homework/src/dto/general"
)

type IUserService interface {
	Register(_user domain.User) (general.Token, error)
}
