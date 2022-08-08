package entity_user

import (
	"property-finder-go-bootcamp-homework/pkg/string_helper"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//UserInfo struct contains all information about user except id
type UserInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//newUserInfo creates new UserInfo struct and returns it with filled fields
func NewUserInfo(firstname, lastname, email, password string) *UserInfo {
	userInfo := new(UserInfo)
	userInfo.SetUserInfoFirstname(firstname)
	userInfo.SetUserInfoLastname(lastname)
	userInfo.SetUserInfoEmail(email)
	userInfo.HashPassword(password)
	return userInfo
}

//setUserInfoFirstname sets firstname to UserInfo struct  with upper case first letter
func (userInfo *UserInfo) SetUserInfoFirstname(Firstname string) {
	userInfo.Firstname = string_helper.UpperCaseFirstLetters(Firstname)
}

//setUserInfoLastname sets lastname to UserInfo struct with upper case first letter
func (userInfo *UserInfo) SetUserInfoLastname(Lastname string) {
	userInfo.Lastname = string_helper.UpperCaseFirstLetters(Lastname)
}

//setUserInfoEmail sets email to UserInfo struct with lower case
func (userInfo *UserInfo) SetUserInfoEmail(Email string) {
	userInfo.Email = strings.ToLower(strings.ReplaceAll(Email, " ", ""))
}

//setUserInfoPassword sets password to UserInfo struct with hashed password
func (userInfo *UserInfo) SetUserInfoPassword(Password string) {
	userInfo.Password = Password
}

//hashPassword hashes password to UserInfo struct
func (userInfo *UserInfo) HashPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	userInfo.SetUserInfoPassword(string(bytes))
}

//checkPassword checks if password is correct
func (userInfo *UserInfo) ComparePasswords(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)) == nil
}
