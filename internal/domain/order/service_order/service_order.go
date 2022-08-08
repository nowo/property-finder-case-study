package service_order

import (
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
)

//OrderService is a struct that implements IServiceOrder interface.
type OrderService struct {
	OrderRepo repository_order.IOrderRepository
	CartRepo  repository_cart.ICartRepository
}

//NewOrderService is a function that returns a new instance of OrderService.
func New(cartRepo repository_cart.ICartRepository, orderRepo repository_order.IOrderRepository) IServiceOrder {
	return &OrderService{
		OrderRepo: orderRepo,
		CartRepo:  cartRepo,
	}
}

//CreateOrder is a function that creates a new order with the given userID, totalPrice and vatOfCart.
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

//GetOrderByUserID is a function that returns all orders of the given userID.
func (s *OrderService) GetOrderByUserID(userID uint) ([]order.Order, error) {
	return s.OrderRepo.GetOrderByUserID(userID)
}
