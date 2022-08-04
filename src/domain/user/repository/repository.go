package repository

import (
	"property-finder-go-bootcamp-homework/src/config/messages"
	"property-finder-go-bootcamp-homework/src/database/postgres"
	domain "property-finder-go-bootcamp-homework/src/domain/user"
	"property-finder-go-bootcamp-homework/src/domain/user/entity"
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
	if err != nil {
		return false, nil
	}
	return true, messages.EMAIL_ALREADY_EXIST
}

func (r *Repository) GetUserInfoByEmail(email string) (entity.UserInfo, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	user := new(domain.User)
	response := db.Table("users").Where("email = ?", email).First(user)

	if response.Error != nil {
		return entity.UserInfo{}, messages.DATABASE_OPERATION_FAILED
	}

	return user.UserInfo, nil
}

func (r *Repository) Create(user domain.User) (domain.User, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&user).Error; err != nil {
		return domain.User{}, messages.DATABASE_OPERATION_FAILED
	}

	return user, nil
}
