package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config/postgres_config"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/user"
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
	DB.AutoMigrate(&user.User{}, &product.Product{}, &cart.Cart{})
}
