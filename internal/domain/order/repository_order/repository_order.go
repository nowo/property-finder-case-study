package repository_order

import (
	"errors"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"time"
)

// OrderRepository is a struct that implements IOrderRepository interface.
type OrderRepository struct {
}

// CreateOrder creates a new order.
func (repository *OrderRepository) CreateOrder(newOrder order.Order) (order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	if err := db.Create(&newOrder).Error; err != nil {
		return order.Order{}, err
	}
	return newOrder, nil
}

// GetOrderByUserID returns all orders of a user from database.
func (repository *OrderRepository) GetOrderByUserID(userID uint) ([]order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	var orders []order.Order
	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orders, messages.NO_RECORD_FOUND
		}
		return nil, err
	}
	return orders, nil
}

// GetOrderFromLastMonth returns all orders from last month from database.
func (repository *OrderRepository) GetOrderFromLastMonth(userID uint) ([]order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	var orders []order.Order
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	if err := db.Where("user_id = ?", userID).Where("created_at BETWEEN ? AND ?", lastMonth, now).Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orders, messages.NO_RECORD_FOUND
		}
		return nil, err
	}
	return orders, nil
}
