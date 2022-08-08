package service_product

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

// IServiceProduct interface contains all methods that are required to implement by service_product.
type IServiceProduct interface {
	GetAll() ([]domain.Product, error)
	GetByID(id uint) (domain.Product, error)
}
