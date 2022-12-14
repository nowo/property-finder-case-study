package repository_user

import (
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/domain/user"
	"property-finder-go-bootcamp-homework/pkg/errors"
)

// IUserRepository interface contains all methods that are required to implement by repository_user.
type IUserRepository interface {
	GetUserInfoByEmail(email string) (user.User, error)
	Create(newUser user.User) (user.User, error)
	CheckEmailExists(email string) (bool, error)
}

//UserRepository is a struct that implements IUserRepository interface
type UserRepository struct {
}

//CheckEmailExists checks if email exists in database
func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	user := new(user.User)
	err := db.Table("users").Where("email = ?", email).First(user)

	if err.RowsAffected == 0 {
		return false, nil
	}
	return true, errors.NewEmailAlreadyExist(email)
}

//GetUserInfoByEmail returns user by email
func (r *UserRepository) GetUserInfoByEmail(email string) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	newUser := new(user.User)
	response := db.Table("users").Where("email = ?", email).First(newUser)
	if response.Error != nil {
		return user.User{}, response.Error
	}

	return *newUser, nil
}

//Create creates a new user
func (r *UserRepository) Create(newUser user.User) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&newUser).Error; err != nil {
		return user.User{}, err
	}

	return newUser, nil
}
