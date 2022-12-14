package user

import (
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"

	"gorm.io/gorm"
)

//Our main aggregate user
type User struct {
	gorm.Model
	UserInfo entity_user.UserInfo `json:"user_info" gorm:"embedded;embedded_prefix:user_info_"`
}
