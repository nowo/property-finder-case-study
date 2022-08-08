package repository_user

import (
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

// IUserRepository interface contains all methods that are required to implement by repository_user.
type IUserRepository interface {
	GetUserInfoByEmail(email string) (user.User, error)
	Create(newUser user.User) (user.User, error)
	CheckEmailExists(email string) (bool, error)
}
