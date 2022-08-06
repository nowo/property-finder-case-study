package repository_order

import (
	"fmt"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"time"
)

type OrderRepository struct {
}

func New() IRepositoryOrder {
	return &OrderRepository{}
}

func (repository *OrderRepository) CreateOrder(newOrder order.Order) (order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	if err := db.Create(&newOrder).Error; err != nil {
		return order.Order{}, err
	}

	return newOrder, nil
}
func (repository *OrderRepository) GetOrderByUserID(userID uint) ([]order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	fmt.Println("userID: ", userID)
	var orders []order.Order
	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		fmt.Println("database faili")
		return nil, messages.DATABASE_OPERATION_FAILED
	}
	return orders, nil
}
func (repository *OrderRepository) GetOrderFromLastMonth(userID uint) ([]order.Order, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)
	var orders []order.Order

	//get time now
	now := time.Now()
	//get time last month
	lastMonth := now.AddDate(0, -1, 0)

	if err := db.Where("user_id = ?", userID).Where("created_at BETWEEN ? AND ?", lastMonth, now).Find(&orders).Error; err != nil {
		return nil, messages.DATABASE_OPERATION_FAILED
	}
	return orders, nil
}
