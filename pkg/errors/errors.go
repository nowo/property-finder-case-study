package errors

type CustomError struct {
	Message string `json:"message"`
}

func (c *CustomError) Error() string {
	return c.Message
}

func NewEmailAlreadyExist(email string) error {
	return &CustomError{Message: "Email " + email + " already exist"}
}
