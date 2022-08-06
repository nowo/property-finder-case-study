package repository_product

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

type IProductRepository interface {
	ShowAllProducts() ([]domain.Product, error)
	GetProductByID(id uint) (domain.Product, error)
}
