package entity

import (
	"property-finder-go-bootcamp-homework/pkg/str"
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

func (userInfo *UserInfo) GetUserInfoFirstname() string {
	return userInfo.Firstname
}

func (userInfo *UserInfo) GetUserInfoLastname() string {
	return userInfo.Lastname
}

func (userInfo *UserInfo) GetUserInfoEmail() string {
	return userInfo.Email
}

func (userInfo *UserInfo) GetUserInfoPassword() string {
	return userInfo.Password
}

func (userInfo *UserInfo) SetUserInfoFirstname(Firstname string) {
	userInfo.Firstname = str.UpperCaseFirstLetters(Firstname)
}

func (userInfo *UserInfo) SetUserInfoLastname(Lastname string) {
	userInfo.Lastname = str.UpperCaseFirstLetters(Lastname)
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
