package service_order

import (
	"fmt"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
)

type OrderService struct {
	OrderRepo repository_order.IRepositoryOrder
	CartRepo  repository_cart.ICartRepository
}

func New() IOrderService {
	return &OrderService{
		OrderRepo: repository_order.New(),
		CartRepo:  repository_cart.New(),
	}
}

func (s *OrderService) CreateOrder(userID uint, totalPrice, vatOfCart float64) error {

	newOrder := entity_order.NewOrderInfo(userID, totalPrice, vatOfCart)
	createResponse, orderCreateErr := s.OrderRepo.CreateOrder(order.Order{
		OrderInfo: *newOrder,
	})
	if orderCreateErr != nil {
		return orderCreateErr
	}
	cartCompleteError := s.CartRepo.Complete(userID, createResponse.ID)
	if cartCompleteError != nil {
		return cartCompleteError
	}
	fmt.Println(s.OrderRepo.GetOrderByUserID(userID))
	fmt.Println(s.OrderRepo.GetOrderFromLastMonth(userID))
	return nil
}

func (s *OrderService) GetOrderByUserID(userID uint) ([]order.Order, error) {
	return s.OrderRepo.GetOrderByUserID(userID)
}
