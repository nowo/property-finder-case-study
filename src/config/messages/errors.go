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
)