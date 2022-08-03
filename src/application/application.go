package application

import (
	"fmt"
	"property-finder-go-bootcamp-homework/src/config/postgres_config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	fmt.Println("Application started!!!")
	fmt.Println(postgres_config.URI)
	dsn := postgres_config.URI
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	fmt.Println("connection succesfully")

}
