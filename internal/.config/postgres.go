package _config

import (
	"fmt"
	"os"
)

var (
	HOST          = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	USER          = os.Getenv("POSTGRES_USER")
	PASSWORD      = os.Getenv("POSTGRES_PASS")
	DBNAME        = os.Getenv("POSTGRES_DB")
	URI           = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", HOST, USER, PASSWORD, DBNAME, POSTGRES_PORT)
)
