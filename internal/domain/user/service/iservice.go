package service

import (
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

type IUserService interface {
	Register(_user domain.User) (general.Token, error)
	Login(_dto auth.LoginRequest) (general.Token, error)
}
