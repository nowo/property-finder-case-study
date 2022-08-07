package _config

import (
	"fmt"
	"os"
)

var (
	PORT = fmt.Sprintf(":%v", os.Getenv("PORT"))
)

const (
	AUTH_ENDPOINT     = "/auth"
	CART_ENDPOINT     = "/cart"
	PRODUCT_ENDPOINT  = "/product"
	ORDER_ENDPOINT    = "/order"
	REGISTER_ENDPOINT = "/register"
	LOGIN_ENDPOINT    = "/login"
	DELETE_ENDPOINT   = "/delete"
	CREATE_ENDPOINT   = "/create"
	LIST_ENDPOINT     = "/list"
	API_VERSION       = "/api/v1"
	EMPTY             = "/"
)
