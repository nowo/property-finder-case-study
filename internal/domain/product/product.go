package product

import (
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductInfo entity_product.ProductInfo `json:"product_info" gorm:"embedded;embedded_prefix:product_info_"`
}
