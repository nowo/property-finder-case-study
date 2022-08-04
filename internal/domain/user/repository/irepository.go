package repository

import (
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

type IRepository interface {
	GetUserInfoByEmail(email string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	CheckEmailExists(email string) (bool, error)
}
