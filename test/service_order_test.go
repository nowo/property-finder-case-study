package test

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/internal/domain/order/service_order"

	"property-finder-go-bootcamp-homework/test/mocks"
	"testing"
)

func Test_GET_ORDER_BY_ID(t *testing.T) {
	Convey("Get Order By ID Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		orderService := service_order.New(mockCartRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		Convey("User Can Go", func() {
			order, err := orderService.GetOrderByUserID(uint(1))
			So(err, ShouldBeNil)
			So(order, ShouldNotBeNil)
		})
	})
}

func Test_ORDER_NOT_FOUND(t *testing.T) {
	Convey("Order Not Found Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		orderService := service_order.New(mockCartRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			orders, err := orderService.GetOrderByUserID(uint(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
			So(orders, ShouldResemble, []order.Order{})
		})
	})
}

func Test_CREATE_ORDER(t *testing.T) {
	Convey("Create Order Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		orderService := service_order.New(mockCartRepository, mockOrderRepository)
		newOrderInfo := entity_order.NewOrderInfo(uint(1), 1, 1)
		mockOrderRepository.EXPECT().CreateOrder(order.Order{
			OrderInfo: *newOrderInfo,
		}).Return(order.Order{}, nil)
		mockCartRepository.EXPECT().Complete(uint(1), uint(0)).Return(nil)
		Convey("User Can Go", func() {
			err := orderService.CreateOrder(uint(1), float64(1), float64(1))
			So(err, ShouldBeNil)
		})
	})
}

func Test_CREATE_ORDER_FAIL(t *testing.T) {
	Convey("Create Order Fail Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		orderService := service_order.New(mockCartRepository, mockOrderRepository)
		newOrderInfo := entity_order.NewOrderInfo(uint(1), 1, 1)
		mockOrderRepository.EXPECT().CreateOrder(order.Order{
			OrderInfo: *newOrderInfo,
		}).Return(order.Order{}, gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			err := orderService.CreateOrder(uint(1), float64(1), float64(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
		})
	})
}
