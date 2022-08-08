package repository_order

import "property-finder-go-bootcamp-homework/internal/domain/order"

// IOrderRepository interface contains all methods that are required to implement by repository_order.
type IOrderRepository interface {
	CreateOrder(newOrder order.Order) (order.Order, error)
	GetOrderByUserID(userID uint) ([]order.Order, error)
	GetOrderFromLastMonth(userID uint) ([]order.Order, error)
}
