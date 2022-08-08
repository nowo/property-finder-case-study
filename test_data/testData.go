package test_data

import (
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"
	"property-finder-go-bootcamp-homework/internal/domain/user/entity_user"
)

var ProductList = []product.Product{
	{
		Model: gorm.Model{
			ID: 1,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product",
			Price: 100,
			Vat:   8,
		}},
	{
		Model: gorm.Model{
			ID: 1,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product",
			Price: 100,
			Vat:   8,
		}},
	{
		Model: gorm.Model{
			ID: 1,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product",
			Price: 100,
			Vat:   8,
		}},
	{
		Model: gorm.Model{
			ID: 1,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product",
			Price: 100,
			Vat:   8,
		}},
	{
		Model: gorm.Model{
			ID: 2,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product2",
			Price: 200,
			Vat:   1,
		}},
	{
		Model: gorm.Model{
			ID: 3,
		},
		ProductInfo: entity_product.ProductInfo{
			Name:  "Product3",
			Price: 500,
			Vat:   18,
		}},
}

var ValidRequestBody = entity_user.UserInfo{
	Firstname: "Erdal",
	Lastname:  "Cinar",
	Email:     "erdalburakcinar@hotmail.com",
	Password:  "123456789",
}
