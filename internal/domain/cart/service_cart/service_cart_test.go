package service_cart

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/order"
	"property-finder-go-bootcamp-homework/internal/domain/order/entity_order"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"
	"property-finder-go-bootcamp-homework/test_data"
	"property-finder-go-bootcamp-homework/test_data/mocks"
	"testing"
)

func Test_PRODUCT_NOT_FOUND(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(product.Product{}, nil)
		Convey("User Can Go", func() {
			err := cartService.AddToCart(uint(1), 1)
			So(err, ShouldResemble, messages.PRODUCT_NOT_FOUND)
		})
	})
}

func Test_CART_ADDED(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		newCart := cart.NewCart(uint(0), uint(0))
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(0)).Return(product.Product{
			ProductInfo: entity_product.ProductInfo{
				Name:     "Product",
				Quantity: 1,
			},
		}, nil)
		mockCartRepository.EXPECT().Create(*newCart).Return(nil)
		Convey("User Can Go", func() {
			err := cartService.AddToCart(uint(0), 0)
			So(err, ShouldBeNil)
		})
	})
}

func Test_CART_NOT_FOUND(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().GetCartsByUserID(uint(1)).Return([]cart.Cart{}, gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			products, err := cartService.GetCartByUserID(uint(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
			So(products, ShouldBeNil)
		})
	})
}

func Test_CART_DELETED(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().Delete(uint(1), uint(1)).Return(nil)
		Convey("User Can Go", func() {
			err := cartService.DeleteFromCart(uint(1), 1)
			So(err, ShouldBeNil)
		})
	})
}

func Test_CartDeleteFailed(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().Delete(uint(1), uint(1)).Return(gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			err := cartService.DeleteFromCart(uint(1), 1)
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
		})
	})
}

func Test_CalculatePriceNoDiscount(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		products := []product.Product{
			{ProductInfo: entity_product.ProductInfo{Name: "Product", Quantity: 1, Price: 100, Vat: 18}},
		}
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{}, nil)

		Convey("User Can Go", func() {
			price, vatOfCart := cartService.CalculatePrice(products, uint(1))
			So(vatOfCart, ShouldEqual, 18)
			So(price, ShouldEqual, 118)
		})
	})
}

func Test_CalculatePiceApplyAfterThreeProductDiscount(t *testing.T) {
	Convey("Given user has same product more than three", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{}, nil)
		Convey("Then apply discount after 3 product", func() {
			price, vatOfCart := cartService.CalculatePrice(test_data.ProductList, uint(1))
			So(vatOfCart, ShouldEqual, 122.4)
			So(price, ShouldEqual, 1202.4)
		})
	})
}

func Test_CalculatePriceApplyMonthlyDiscount(t *testing.T) {
	Convey("Given User Passed Given Amount", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
		}, nil)
		Convey("Apply monthly discount price", func() {
			price, vatOfCart := cartService.CalculatePrice(test_data.ProductList, uint(1))
			So(vatOfCart, ShouldEqual, 111.6)
			So(price, ShouldEqual, 1101.6)
		})
	})
}

func Test_CalculatePriceApplyForthOrderDiscount(t *testing.T) {
	Convey("Given User Passed Given Amount", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
		}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{}, nil)
		Convey("Apply forth order discount price", func() {
			price, vatOfCart := cartService.CalculatePrice(test_data.ProductList, uint(1))
			So(vatOfCart, ShouldEqual, 107.3)
			So(price, ShouldEqual, 1092.3)
		})
	})
}
