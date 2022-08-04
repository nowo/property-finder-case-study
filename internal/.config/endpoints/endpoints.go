package endpoints

import (
	"fmt"
	"os"
)

var (
	PORT = fmt.Sprintf(":%v", os.Getenv("PORT"))
)

const (
	AUTH_ENDPOINT         = "/auth"
	REGISTER_ENDPOINT     = "/register"
	LOGIN_ENDPOINT        = "/login"
	GET_PRODUCTS_ENDPOINT = "/products"
	GET_PRODUCT           = "/product"
)
