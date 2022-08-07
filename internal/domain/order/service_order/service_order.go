package service_order

import (
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
)

type OrderService struct {
	OrderRepo repository_order.IRepositoryOrder
	CartRepo  repository_cart.ICartRepository
}

func New(cartRepo repository_cart.ICartRepository, orderRepo repository_order.IRepositoryOrder) IOrderService {
	return &OrderService{
		OrderRepo: orderRepo,
		CartRepo:  cartRepo,
	}
}

func (s *OrderService) CreateOrder(userID uint, totalPrice, vatOfCart float64) error {
	newOrderInfo := entity_order.NewOrderInfo(userID, totalPrice, vatOfCart)
	createResponse, err := s.OrderRepo.CreateOrder(order.Order{
		OrderInfo: *newOrderInfo,
	})
	if err != nil {
		return err
	}
	err = s.CartRepo.Complete(userID, createResponse.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) GetOrderByUserID(userID uint) ([]order.Order, error) {
	return s.OrderRepo.GetOrderByUserID(userID)
}
