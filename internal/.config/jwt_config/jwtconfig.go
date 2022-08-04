package jwt_config

import (
	"os"
	"time"
)

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
var TOKEN_EXPIRATION_TIME = time.Now().AddDate(0, 3, 0) // 3 Months
