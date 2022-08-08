package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/user"
)

// ConnectDB is used to connect to database
func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(_config.URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

// CloseDB is used to close database connection
func Disconnect(DB *gorm.DB) {
	connection, err := DB.DB()
	if err != nil {
		panic(err)
	}
	connection.Close()
}

// Migration is used to create tables in database
func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&user.User{}, &product.Product{}, &cart.Cart{}, &order.Order{})
}
