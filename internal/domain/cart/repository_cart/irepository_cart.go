package repository_cart

import (
	"property-finder-go-bootcamp-homework/internal/domain/cart"
)

type ICartRepository interface {
	GetCartInfoByUserID(userID uint) ([]cart.Cart, error)
	Create(newCart cart.Cart) error
	CountByProductID(productID uint) (int64, error)
	IsAmountExceedByMonth(userID uint) (bool, error)
	Delete(userID, productID uint) error
}
