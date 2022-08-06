package repository_user

import (
	"fmt"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/user"
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
		fmt.Println("Email not exist")
		return false, nil
	}
	fmt.Println("Email already exist!!!")
	return true, messages.EMAIL_ALREADY_EXIST
}

func (r *Repository) GetUserInfoByEmail(email string) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	newUser := new(user.User)
	response := db.Table("users").Where("email = ?", email).First(newUser)

	if response.Error != nil {
		return user.User{}, messages.DATABASE_OPERATION_FAILED
	}

	return *newUser, nil
}

func (r *Repository) Create(newUser user.User) (user.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&newUser).Error; err != nil {
		return user.User{}, messages.DATABASE_OPERATION_FAILED
	}

	return newUser, nil
}
