package repository_user

import (
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/pkg/errors"
)

type Repository struct {
}

func New() IRepository {
	return &Repository{}
}

func (r *Repository) CheckEmailExists(email string) (bool, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	user := new(user.User)
	err := db.Table("users").Where("email = ?", email).First(user)

	if err.RowsAffected == 0 {
		return false, nil
	}
	return true, errors.NewEmailAlreadyExist(email)
}

func (r *Repository) GetUserInfoByEmail(email string) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	newUser := new(user.User)
	response := db.Table("users").Where("email = ?", email).First(newUser)
	if response.Error != nil {
		return user.User{}, response.Error
	}

	return *newUser, nil
}

func (r *Repository) Create(newUser user.User) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&newUser).Error; err != nil {
		return user.User{}, err
	}

	return newUser, nil
}
