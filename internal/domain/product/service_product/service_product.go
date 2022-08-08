package service_product

import (
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
)

// ProductService is a struct that implements the IServiceProduct interface.
type ProductService struct {
	Repo repository_product.IProductRepository
}

// NewProductService is a function that returns a pointer to a new ProductService struct.
func New(repo repository_product.IProductRepository) IServiceProduct {
	return &ProductService{
		Repo: repo,
	}
}

// GetAll is a function that returns all products.
func (s *ProductService) GetAll() ([]domain.Product, error) {
	products, err := s.Repo.ShowAllProducts()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

// GetByID is a function that returns a product by its ID.
func (s *ProductService) GetByID(id uint) (domain.Product, error) {
	product, err := s.Repo.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
