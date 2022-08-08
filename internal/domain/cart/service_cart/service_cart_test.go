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
	"property-finder-go-bootcamp-homework/test_data/repository_mocks"
	"testing"
)

func Test_ProductNotFound(t *testing.T) {
	Convey("Given that i tried to add unable product to cart ", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(product.Product{}, nil)
		Convey("Then i get product not enough quantity  error", func() {
			err := cartService.AddToCart(uint(1), 1)
			So(err, ShouldResemble, messages.NOT_ENOUGH_QUANTITY)
		})
	})
}

func Test_AddProductToCart(t *testing.T) {
	Convey("Given that i add valid product to cart ", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		newCart := cart.NewCart(uint(1), uint(1))
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(product.Product{
			Model: gorm.Model{
				ID: uint(1),
			},
			ProductInfo: entity_product.ProductInfo{
				Name:     "Product",
				Quantity: 2,
			},
		}, nil)
		mockProductRepository.EXPECT().UpdateProductQuantity(uint(1), 1).Return(nil)
		mockCartRepository.EXPECT().Create(*newCart).Return(nil)
		Convey("Then i product added into my cart", func() {
			err := cartService.AddToCart(uint(1), 1)
			So(err, ShouldBeNil)
		})
	})
}

func Test_UserNotFound(t *testing.T) {
	Convey("Given that i tried to add product to cart with invalid user", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockCartRepository.EXPECT().GetCartsByUserID(uint(1)).Return([]cart.Cart{}, gorm.ErrRecordNotFound)
		Convey("Then i get record not found error", func() {
			products, err := cartService.GetCartByUserID(uint(1))
			So(err, ShouldResemble, gorm.ErrRecordNotFound)
			So(products, ShouldBeNil)
		})
	})
}

func Test_DeleteProductFromCart(t *testing.T) {
	Convey("Given that i tried to delete product from cart", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(product.Product{}, nil)
		mockProductRepository.EXPECT().UpdateProductQuantity(uint(1), 1).Return(nil)
		mockCartRepository.EXPECT().Delete(uint(1), uint(1)).Return(nil)
		Convey("Then product deleted from cart", func() {
			err := cartService.DeleteFromCart(uint(1), 1)
			So(err, ShouldBeNil)
		})
	})
}

func Test_CartDeleteFailed(t *testing.T) {
	Convey("Given that i tried to delete invalid product from cart", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(product.Product{}, nil)
		mockProductRepository.EXPECT().UpdateProductQuantity(uint(1), 1).Return(nil)
		mockCartRepository.EXPECT().Delete(uint(1), uint(1)).Return(gorm.ErrRecordNotFound)
		Convey("Then i get record not found error", func() {
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
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
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
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
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
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
		}, nil)
		Convey("Then apply monthly discount price", func() {
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
		mockOrderRepository := repository_mocks.NewMockIOrderRepository(mockCtrl)
		mockCartRepository := repository_mocks.NewMockICartRepository(mockCtrl)
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		cartService := New(mockCartRepository, mockProductRepository, mockOrderRepository)
		mockOrderRepository.EXPECT().GetOrderByUserID(uint(1)).Return([]order.Order{
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
			{OrderInfo: entity_order.OrderInfo{UserID: 1, TotalPrice: 50000}},
		}, nil)
		mockOrderRepository.EXPECT().GetOrderFromLastMonth(uint(1)).Return([]order.Order{}, nil)
		Convey("Then apply forth order discount price", func() {
			price, vatOfCart := cartService.CalculatePrice(test_data.ProductList, uint(1))
			So(vatOfCart, ShouldEqual, 107.3)
			So(price, ShouldEqual, 1092.3)
		})
	})
}
