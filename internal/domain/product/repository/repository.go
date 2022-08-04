package repository

import (
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
)

type ProductRepository struct {
}

func New() IProductRepository {
	return &ProductRepository{}
}
func (r *ProductRepository) ShowAllProducts() ([]domain.Product, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	var products []domain.Product
	response := db.Table("products").Where("quantity > 0").Find(&products)

	if response.Error != nil {
		return nil, messages.DATABASE_OPERATION_FAILED
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (domain.Product, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	var product domain.Product
	response := db.Table("products").Where("id = ?", id).First(&product)

	if response.Error != nil {
		return domain.Product{}, messages.PRODUCT_NOT_FOUND
	}

	return product, nil
}
