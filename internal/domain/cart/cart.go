package cart

import (
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/cart/entity_cart"
)

//Our main aggregate cart
type Cart struct {
	gorm.Model
	CartInfo entity_cart.CartInfo `json:"cart_info" gorm:"embedded;embedded_prefix:cart_info_"`
}

func NewCart(userID uint, productID uint) *Cart {
	return &Cart{
		CartInfo: *entity_cart.NewCartInfo(userID, productID, 0, false),
	}
}
