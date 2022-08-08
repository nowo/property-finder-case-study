package service_product

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	domain "property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/entity_product"
	"property-finder-go-bootcamp-homework/test_data/repository_mocks"
	"testing"
)

func Test_GetAllProducts(t *testing.T) {
	Convey("Given that i tried to get all products", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().ShowAllProducts().Return([]domain.Product{}, nil)
		Convey("Then i get all products", func() {
			products, err := productService.GetAll()
			So(err, ShouldBeNil)
			So(products, ShouldNotBeNil)
		})
	})
}

func Test_GetProductByID(t *testing.T) {
	Convey("Given that i tried to get spesific product with id", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{
			gorm.Model{
				ID: 1,
			},
			entity_product.ProductInfo{
				Name: "Product 1",
			},
		}, nil)
		Convey("Then i get spesific product", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldBeNil)
			So(product, ShouldNotBeNil)
		})
	})
}

func TestProductNotFound(t *testing.T) {
	Convey("Given that i searched invalid product", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{}, gorm.ErrRecordNotFound)
		Convey("Then i get error", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldNotBeNil)
			So(product, ShouldResemble, domain.Product{})
		})
	})
}

func Test_Product_Service_Error(t *testing.T) {
	Convey("Given that databased bugged when searching product", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockProductRepository := repository_mocks.NewMockIProductRepository(mockCtrl)
		productService := New(mockProductRepository)
		mockProductRepository.EXPECT().GetProductByID(uint(1)).Return(domain.Product{}, messages.DATABASE_OPERATION_FAILED)
		Convey("Then i get database operation failed error", func() {
			product, err := productService.GetByID(1)
			So(err, ShouldNotBeNil)
			So(product, ShouldResemble, domain.Product{})
		})
	})
}
