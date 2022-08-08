package repository_product

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

// IRepositoryProduct interface contains all methods that are required to implement by repository_product.
type IProductRepository interface {
	ShowAllProducts() ([]domain.Product, error)
	GetProductByID(id uint) (domain.Product, error)
	UpdateProductQuantity(id uint, quantity int) error
}
