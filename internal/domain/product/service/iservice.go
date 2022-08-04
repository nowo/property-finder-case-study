package service

import domain "property-finder-go-bootcamp-homework/internal/domain/product"

type IProductService interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
}
