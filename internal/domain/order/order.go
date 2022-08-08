package order

import (
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
)

//Our main aggregate order
type Order struct {
	gorm.Model
	OrderInfo entity_order.OrderInfo `json:"order_info" gorm:"embedded;embedded_prefix:order_info_"`
}
