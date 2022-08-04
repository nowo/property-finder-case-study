package application

import (
	"fmt"
	"property-finder-go-bootcamp-homework/src/database/postgres"
	"property-finder-go-bootcamp-homework/src/fiber/router"
)

func Start() {
	fmt.Println("Application started")
	postgres.Migration(postgres.ConnectDB())
	router.Router()
}
