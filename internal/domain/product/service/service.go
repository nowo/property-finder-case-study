package service

import (
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository"
)

type ProductService struct {
	Repo repository.IProductRepository
}

func New() IProductService {
	return &ProductService{
		Repo: repository.New(),
	}
}
func (s *ProductService) GetAll() ([]domain.Product, error) {

	return s.Repo.ShowAllProducts()
}

func (s *ProductService) GetByID(id int) (domain.Product, error) {
	return s.Repo.GetProductByID(id)
}
