package endpoints

import (
	"fmt"
	"os"
)

var (
	PORT = fmt.Sprintf(":%v", os.Getenv("PORT"))
)

const (
	AUTH_ENDPOINT    = "/auth"
	CART_ENDPOINT    = "/cart"
	PRODUCT_ENDPOINT = "/product"

	REGISTER_ENDPOINT = "/register"
	LOGIN_ENDPOINT    = "/login"

	PRODUCTS_ENDPOINT = "/products"
	DELETE_ENDPOINT   = "/delete"

	API_VERSION = "/api/v1"
	EMPTY       = "/"
	PARAMS_ID   = "/:id"
)
