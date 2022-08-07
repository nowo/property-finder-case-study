package test

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/service_cart"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"
	"property-finder-go-bootcamp-homework/test/mocks"
	"testing"
)

func Test_PRODUCT_NOT_FOUND(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := service_cart.New(mockCartRepository, mockProductRepository, mockOrderRepository)
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
		cartService := service_cart.New(mockCartRepository, mockProductRepository, mockOrderRepository)
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

func Test_CART_DELETED(t *testing.T) {
	Convey("Create Cart Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := mocks.NewMockIRepositoryOrder(mockCtrl)
		mockCartRepository := mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		cartService := service_cart.New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().Delete(uint(1), uint(1)).Return(nil)
		Convey("User Can Go", func() {
			err := cartService.DeleteFromCart(uint(1), 1)
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
		cartService := service_cart.New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().GetCartsByUserID(uint(1)).Return([]cart.Cart{}, gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			products, err := cartService.GetCartByUserID(uint(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
			So(products, ShouldBeNil)
		})
	})
}
