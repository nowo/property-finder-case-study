package service_user

import (
	"property-finder-go-bootcamp-homework/dto/auth"
	"property-finder-go-bootcamp-homework/dto/general"
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

type IUserService interface {
	Register(_user user.User) (general.Token, error)
	Login(_dto auth.LoginRequest) (general.Token, error)
}
