package repository

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

type IProductRepository interface {
	ShowAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
}
