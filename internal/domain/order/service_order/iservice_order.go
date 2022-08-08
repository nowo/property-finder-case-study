package service_order

import "property-finder-go-bootcamp-homework/internal/domain/order"

// IServiceOrder interface contains all methods that are required to implement by service_order.
type IServiceOrder interface {
	CreateOrder(userID uint, totalPrice, vatOfCart float64) error
	GetOrderByUserID(userID uint) ([]order.Order, error)
}
