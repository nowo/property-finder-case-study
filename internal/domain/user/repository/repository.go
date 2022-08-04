package repository

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
	user := new(domain.User)
	err := db.Table("users").Where("email = ?", email).First(user)

	if err.RowsAffected == 0 {
		fmt.Println("Email not exist")
		return false, nil
	}
	fmt.Println("Email already exist!!!")
	return true, messages.EMAIL_ALREADY_EXIST
}

func (r *Repository) GetUserInfoByEmail(email string) (domain.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	user := new(domain.User)
	response := db.Table("users").Where("email = ?", email).First(user)

	if response.Error != nil {
		return domain.User{}, messages.DATABASE_OPERATION_FAILED
	}

	return *user, nil
}

func (r *Repository) Create(user domain.User) (domain.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&user).Error; err != nil {
		return domain.User{}, messages.DATABASE_OPERATION_FAILED
	}

	return user, nil
}
