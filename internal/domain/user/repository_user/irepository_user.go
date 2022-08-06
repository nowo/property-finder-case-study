package repository_user

import (
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

type IRepository interface {
	GetUserInfoByEmail(email string) (user.User, error)
	Create(user user.User) (user.User, error)
	CheckEmailExists(email string) (bool, error)
}
