package domain

import (
	"property-finder-go-bootcamp-homework/src/domain/user/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserInfo entity.UserInfo `json:"user_info" gorm:"embedded;embedded_prefix:user_info_"`
}

func (user *User) GetUserInfo() entity.UserInfo {
	return user.UserInfo
}

func (user *User) SetUserInfo(UserInfo entity.UserInfo) {
	user.UserInfo = UserInfo
}
