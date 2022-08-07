package entity_user

import (
	"property-finder-go-bootcamp-homework/pkg/string_helper"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUserInfo(firstname, lastname, email, password string) *UserInfo {
	userInfo := new(UserInfo)
	userInfo.SetUserInfoFirstname(firstname)
	userInfo.SetUserInfoLastname(lastname)
	userInfo.SetUserInfoEmail(email)
	userInfo.HashPassword(password)
	return userInfo
}

func (userInfo *UserInfo) SetUserInfoFirstname(Firstname string) {
	userInfo.Firstname = string_helper.UpperCaseFirstLetters(Firstname)
}

func (userInfo *UserInfo) SetUserInfoLastname(Lastname string) {
	userInfo.Lastname = string_helper.UpperCaseFirstLetters(Lastname)
}

func (userInfo *UserInfo) SetUserInfoEmail(Email string) {
	userInfo.Email = strings.ToLower(strings.ReplaceAll(Email, " ", ""))
}

func (userInfo *UserInfo) SetUserInfoPassword(Password string) {
	userInfo.Password = Password
}

func (userInfo *UserInfo) HashPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	userInfo.SetUserInfoPassword(string(bytes))
}

func (userInfo *UserInfo) ComparePasswords(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)) == nil
}
