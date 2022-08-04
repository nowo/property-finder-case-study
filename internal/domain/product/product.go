package domain

import (
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity"
)

type Product struct {
	gorm.Model
	ProductInfo entity.ProductInfo `json:"product_info" gorm:"embedded;embedded_prefix:product_info_"`
}

func (product *Product) GetProductInfo() *entity.ProductInfo {
	return &product.ProductInfo
}
