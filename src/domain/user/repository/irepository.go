package repository

import (
	domain "property-finder-go-bootcamp-homework/src/domain/user"
	"property-finder-go-bootcamp-homework/src/domain/user/entity"
)

type IRepository interface {
	GetUserInfoByEmail(email string) (entity.UserInfo, error)
	Create(user domain.User) (domain.User, error)
	CheckEmailExists(email string) (bool, error)
}
