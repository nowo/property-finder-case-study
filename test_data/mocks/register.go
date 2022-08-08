package mocks

import "property-finder-go-bootcamp-homework/internal/domain/user/entity_user"

var ValidRequestBody = entity_user.UserInfo{
	Firstname: "Erdal",
	Lastname:  "Cinar",
	Email:     "erdalburakcinar@hotmail.com",
	Password:  "123456789",
}
