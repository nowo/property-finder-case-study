package service_product

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

type IProductService interface {
	GetAll() ([]domain.Product, error)
	GetByID(id uint) (domain.Product, error)
}
