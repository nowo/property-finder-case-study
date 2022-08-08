package service_order

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/test_data/repository_mocks"
	"testing"
)

func Test_GetOrderByID(t *testing.T) {
	Convey("Given that i tried to get my orders", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		orderService := New(mockCartRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		Convey("Then i get my orders", func() {
			order, err := orderService.GetOrderByUserID(uint(1))
			So(err, ShouldBeNil)
			So(order, ShouldNotBeNil)
		})
	})
}

func Test_OrderNotFound(t *testing.T) {
	Convey("Given that i tried to search not found order", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		orderService := New(mockCartRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, gorm.ErrRecordNotFound)
		Convey("Then i get record not found error", func() {
			orders, err := orderService.GetOrderByUserID(uint(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
			So(orders, ShouldResemble, []order.Order{})
		})
	})
}

func Test_CreateOrder(t *testing.T) {
	Convey("Given that i tried to complete my order", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		orderService := New(mockCartRepository, mockOrderRepository)
		newOrderInfo := entity_order.NewOrderInfo(uint(1), 1, 1)
		mockOrderRepository.EXPECT().CreateOrder(order.Order{
			OrderInfo: *newOrderInfo,
		}).Return(order.Order{}, nil)
		mockCartRepository.EXPECT().Complete(uint(1), uint(0)).Return(nil)
		Convey("Then i completed my order sucessfully", func() {
			err := orderService.CreateOrder(uint(1), float64(1), float64(1))
			So(err, ShouldBeNil)
		})
	})
}

func Test_CreateOrderFailed(t *testing.T) {
	Convey("Given that i tried to complete my order with invalid id ", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		orderService := New(mockCartRepository, mockOrderRepository)
		newOrderInfo := entity_order.NewOrderInfo(uint(1), 1, 1)
		mockOrderRepository.EXPECT().CreateOrder(order.Order{
			OrderInfo: *newOrderInfo,
		}).Return(order.Order{}, gorm.ErrRecordNotFound)
		Convey("Then i get record not found error", func() {
			err := orderService.CreateOrder(uint(1), float64(1), float64(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
		})
	})
}
