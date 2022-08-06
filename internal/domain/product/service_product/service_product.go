package service_product

import (
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
)

type ProductService struct {
	Repo repository_product.IProductRepository
}

func New() IProductService {
	return &ProductService{
		Repo: repository_product.New(),
	}
}
func (s *ProductService) GetAll() ([]domain.Product, error) {
	products, getAllProductsError := s.Repo.ShowAllProducts()
	if getAllProductsError != nil {
		return []domain.Product{}, getAllProductsError
	}
	return products, nil
}

func (s *ProductService) GetByID(id uint) (domain.Product, error) {
	product, getProductError := s.Repo.GetProductByID(id)
	if getProductError != nil {
		return domain.Product{}, getProductError
	}
	return product, nil
}
