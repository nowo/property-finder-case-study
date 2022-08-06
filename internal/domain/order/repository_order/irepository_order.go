package repository_order

import "property-finder-go-bootcamp-homework/internal/domain/order"

type IRepositoryOrder interface {
	CreateOrder(newOrder order.Order) (order.Order, error)
	GetOrderByUserID(userID uint) ([]order.Order, error)
	GetOrderFromLastMonth(userID uint) ([]order.Order, error)
}
