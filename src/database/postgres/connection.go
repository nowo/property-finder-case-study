package postgres

import (
	"property-finder-go-bootcamp-homework/src/config/postgres_config"
	domain "property-finder-go-bootcamp-homework/src/domain/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	DB, err := gorm.Open(postgres.Open(postgres_config.URI), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return DB
}

func Disconnect(DB *gorm.DB) {
	connection, err := DB.DB()
	if err != nil {
		panic(err)
	}
	connection.Close()
}

func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&domain.User{})
}
