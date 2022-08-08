package repository_product

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/product"
)

//ProductRepository is a struct that implements the IProductRepository interface.
type ProductRepository struct {
}

//ShowAllProducts returns all products from the database.
func (r *ProductRepository) ShowAllProducts() ([]product.Product, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	var products []product.Product
	response := db.Table("products").Where("quantity > 0").Find(&products)
	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return []product.Product{}, messages.PRODUCT_NOT_FOUND
		}
		return nil, messages.DATABASE_OPERATION_FAILED
	}
	return products, nil
}

//GetProductByID returns a product by its id from the database.
func (r *ProductRepository) GetProductByID(id uint) (product.Product, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	var newProduct product.Product
	response := db.Table("products").Where("id = ?", id).First(&newProduct)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return product.Product{}, messages.PRODUCT_NOT_FOUND
		}
		log.Fatalf("%+v", response.Error)
		return product.Product{}, messages.DATABASE_OPERATION_FAILED
	}

	return newProduct, nil
}

//UpdateProductQuantity updates the quantity of a product from database. it is called when a product is added to cart or when product deleted from cart.
func (r *ProductRepository) UpdateProductQuantity(id uint, quantity int) error {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	response := db.Table("products").Where("id = ?", id).Update("quantity", quantity)
	if response.Error != nil {
		return response.Error
	}

	return nil
}

// IRepositoryProduct interface contains all methods that are required to implement by repository_product.
type IProductRepository interface {
	ShowAllProducts() ([]product.Product, error)
	GetProductByID(id uint) (product.Product, error)
	UpdateProductQuantity(id uint, quantity int) error
}
