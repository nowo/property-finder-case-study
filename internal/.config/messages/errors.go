package messages

import "errors"

var (
	EMAIL_ALREADY_EXIST       = errors.New("Email already exist")
	DATABASE_OPERATION_FAILED = errors.New("Database operation failed")
	REQUIRED_FIELD            = errors.New("Required field")
	WRONG_LENGHT              = errors.New("Wrong length")
	WRONG_FORMAT              = errors.New("Wrong format")
	BAD_REQUEST               = errors.New("Bad request")
	USER_NOT_FOUND            = errors.New("User not found")
	AUTHORIZED_USER           = errors.New("Authorized user")
	UNAUTHORIZED_USER         = errors.New("Unauthorized user")
	INVALID_PASSWORD          = errors.New("Invalid password")
	PRODUCT_NOT_FOUND         = errors.New("Product not found")
	NOT_ENOUGH_QUANTITY       = errors.New("Not enough quantity")
)
