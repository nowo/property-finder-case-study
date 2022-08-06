package service_cart

import (
	"fmt"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/internal/domain/user/repository_user"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
)

type CartService struct {
	CartRepo    repository_cart.ICartRepository
	UserRepo    repository_user.IRepository
	ProductRepo repository_product.IProductRepository
	jwt         _jwt.JWT
}

func New() ICartService {
	return &CartService{
		CartRepo:    repository_cart.New(),
		UserRepo:    repository_user.New(),
		ProductRepo: repository_product.New(),
		jwt:         *_jwt.New(),
	}
}

// AddToCart implements ICartService
func (c *CartService) AddToCart(userID, productID uint) error {

	newCart := cart.NewCart(userID, productID)

	selectedProduct, err := c.ProductRepo.GetProductByID(productID)
	if err != nil {
		return err
	}
	if selectedProduct.ProductInfo.Quantity == 0 {
		return messages.PRODUCT_NOT_FOUND
	}

	if err := c.CartRepo.Create(*newCart); err != nil {
		return err
	}

	//selectedProduct, productError := c.ProductRepo.GetProductByID(productID)
	//if productError != nil {
	//	return productError
	//}
	//productCount, productCountError := c.CartRepo.CountByProductID(productID)
	//if productCountError != nil {
	//	return productCountError
	//}
	//if productCount > 3 {
	//	price += selectedProduct.ProductInfo.Price*3 + selectedProduct.ProductInfo.Price*float64(productCount-3)*0.8
	//}
	//
	//if isAmounthExceed, isAmounthExceedError := c.CartRepo.IsAmountExceedByMonth(userID); isAmounthExceedError != nil {
	//	return isAmounthExceedError
	//} else if isAmounthExceed {
	//	discount += 10
	//}

	return nil
}

func (c *CartService) DeleteFromCart(userID, productID uint) error {

	return c.CartRepo.Delete(userID, productID)
}

func (c *CartService) GetCartByUserID(userID uint) ([]product.Product, error) {

	var buyedProducts = make([]product.Product, 0)
	cartList, err := c.CartRepo.GetCartInfoByUserID(userID)
	if err != nil {
		return nil, err
	}
	for _, cart := range cartList {
		selectedProduct, err := c.ProductRepo.GetProductByID(cart.ProductID)
		if err != nil {
			return nil, err
		}
		buyedProducts = append(buyedProducts, selectedProduct)
	}
	return buyedProducts, nil
}

func (c *CartService) CalculatePrice(cartList []product.Product) (float64, float64, error) {

	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	for _, product := range cartList {
		vat := (product.ProductInfo.Price * float64(product.ProductInfo.Vat)) / 100
		vatOfCart += vat
		totalPrice += product.ProductInfo.Price + vat
	}
	fmt.Println(totalPrice)
	fmt.Println("discount b", moreThanThreeDisountPrice(cartList))
	return totalPrice, vatOfCart, nil
}

func moreThanThreeDisountPrice(productList []product.Product) float64 {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	productCountById := make(map[product.Product]int)
	for _, product := range productList {
		productCountById[product]++
	}

	for selectedProduct, productCount := range productCountById {
		if productCount > 3 {
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			price := selectedProduct.ProductInfo.Price + vat
			totalPrice += price*3 + price*float64(productCount-3)*0.8
		} else {
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			totalPrice += (selectedProduct.ProductInfo.Price + vat) * float64(productCount)
		}
	}
	return totalPrice
}
