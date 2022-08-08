package service_product

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"
	"property-finder-go-bootcamp-homework/test_data/mocks"
	"testing"
)

func Test_Get_All_Products(t *testing.T) {
	Convey("Get All Products Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().ShowAllProducts().Return([]domain.Product{}, nil)
		Convey("User Can Go", func() {
			products, err := productService.GetAll()
			So(err, ShouldBeNil)
			So(products, ShouldNotBeNil)
		})
	})
}

func Test_Get_Product_By_ID(t *testing.T) {
	Convey("Get Product By ID Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{
			gorm.Model{
				ID: 1,
			},
			entity_product.ProductInfo{
				Name: "Product 1",
			},
		}, nil)
		Convey("User Can Go", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldBeNil)
			So(product, ShouldNotBeNil)
		})
	})
}

func Test_Product_Not_Found(t *testing.T) {
	Convey("Product Not Found Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{}, gorm.ErrRecordNotFound)
		Convey("User Can Go", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldNotBeNil)
			So(product, ShouldResemble, domain.Product{})
		})
	})
}

func Test_Product_Service_Error(t *testing.T) {
	Convey("Product Service Error Test Integration", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{}, messages.DATABASE_OPERATION_FAILED)
		Convey("User Can Go", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldNotBeNil)
			So(product, ShouldResemble, domain.Product{})
		})
	})
}
