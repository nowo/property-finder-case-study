package service_product

import (
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
)

type ProductService struct {
	Repo repository_product.IProductRepository
}

func New(repo repository_product.IProductRepository) IProductService {
	return &ProductService{
		Repo: repo,
	}
}
func (s *ProductService) GetAll() ([]domain.Product, error) {
	products, err := s.Repo.ShowAllProducts()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (s *ProductService) GetByID(id uint) (domain.Product, error) {
	product, err := s.Repo.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
