package application

import (
	"fmt"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/api/router"
)

// Application start by calling this function.
func Start() {
	fmt.Println("Application started!!!")
	postgres.Migration(postgres.ConnectDB())
	router.Router()
}
