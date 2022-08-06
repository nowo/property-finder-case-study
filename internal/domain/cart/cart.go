package cart

import (
	"fmt"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/product"
)

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func NewCart(userID uint, productID uint) *Cart {
	fmt.Println(userID)
	return &Cart{
		UserID:    userID,
		ProductID: productID,
	}
}

type Basket struct {
	Cart       []product.Product `json:"cart"`
	TotalPrice float64           `json:"price"`
	VatOfCart  float64           `json:"vat"`
}
