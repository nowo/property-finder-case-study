package _config

import (
	"os"
	"time"
)

var JWT_SECRETKEY = os.Getenv("JWT_SECRET_KEY")
var TOKEN_EXPIRATION_TIME = time.Now().AddDate(0, 3, 0) // 3 Months
