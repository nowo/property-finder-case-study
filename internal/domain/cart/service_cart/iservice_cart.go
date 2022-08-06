package service_cart

import (
	"property-finder-go-bootcamp-homework/internal/domain/product"
)

type ICartService interface {
	AddToCart(userID, productID uint) error
	DeleteFromCart(userID, productID uint) error
	GetCartByUserID(userID uint) ([]product.Product, error)
	CalculatePrice(cartList []product.Product) (float64, float64, error)
}
