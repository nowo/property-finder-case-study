package service_cart

import (
	"fmt"
	"property-finder-go-bootcamp-homework/internal/.config/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
	"strconv"
)

type CartService struct {
	CartRepo    repository_cart.ICartRepository
	ProductRepo repository_product.IProductRepository
	OrderRepo   repository_order.IRepositoryOrder
	jwt         _jwt.JWT
}

func New() ICartService {
	return &CartService{
		CartRepo:    repository_cart.New(),
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

func (c *CartService) CalculatePrice(cartList []product.Product, userID uint) (float64, float64, error) {

	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	for _, product := range cartList {
		vat := (product.ProductInfo.Price * float64(product.ProductInfo.Vat)) / 100
		vatOfCart += vat
		totalPrice += product.ProductInfo.Price + vat
	}
	//discountAPrice, discountBPrice, discountCPrice, discountAVat, discountBVat, discountCVat := totalPrice, totalPrice, totalPrice, vatOfCart, vatOfCart, vatOfCart
	//
	//if c.isUserDeservedForthOrderDiscount(userID, totalPrice) {
	//	discountAPrice, discountAVat = c.calculateForthOrderDiscount(cartList)
	//}
	//discountBPrice, discountBVat = c.moreThanThreeDisountPrice(cartList)
	////discountCPrice, discountCVat = c.monthlyDiscount(userID, totalPrice, vatOfCart)
	//
	//if discountAPrice < discountBPrice {
	//	if discountAPrice < discountCPrice {
	//		return discountAPrice, discountAVat, nil
	//	}
	//	return discountCPrice, discountCVat, nil
	//}
	//if discountBPrice < discountCPrice {
	//	return discountBPrice, discountBVat, nil
	//}

	return totalPrice, vatOfCart, nil
}

func (c *CartService) moreThanThreeDisountPrice(productList []product.Product) (float64, float64) {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	productCountById := make(map[product.Product]int)
	for _, product := range productList {
		productCountById[product]++
	}

	for selectedProduct, productCount := range productCountById {
		if productCount > 3 {
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat*3 + vat*float64(productCount-3)*0.8
			price := selectedProduct.ProductInfo.Price + vat
			totalPrice += price*3 + price*float64(productCount-3)*0.8
		} else {
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			totalPrice += (selectedProduct.ProductInfo.Price + vat) * float64(productCount)
		}
	}
	return totalPrice, vatOfCart
}

func (c *CartService) MonthlyDiscount(userID uint, price float64, vatOfCart float64) (float64, float64) {
	orders, err := c.OrderRepo.GetOrderFromLastMonth(userID)
	if err != nil {
		return price, 0
	}
	var totalPrice float64 = 0
	for _, order := range orders {
		totalPrice += order.OrderInfo.TotalPrice
	}
	givenAmounth, err := strconv.Atoi(general.GIVEN_AMOUNT)
	if err != nil {
		return price, 0
	}
	if totalPrice < float64(givenAmounth) {
		return price, 0
	}
	return price - (price * 0.1), vatOfCart - (vatOfCart * 0.1)
}

func (c *CartService) IsUserDeservedForthOrderDiscount(userID uint, totalPrice float64) bool {
	fmt.Println("isUserDeservedForthOrderDiscount")
	fmt.Println(userID)
	orders, err := c.OrderRepo.GetOrderByUserID(userID)
	if err != nil {
		fmt.Println("isUserDeservedForthOrderDiscount")
		return false
	}
	orderCount := 0
	givenAmounth, err := strconv.Atoi(general.GIVEN_AMOUNT)
	if err != nil {
		return false
	}
	for _, order := range orders {
		if order.OrderInfo.TotalPrice > float64(givenAmounth) {
			orderCount++
		}
	}

	if orderCount%4 != 3 {
		return false
	}
	if totalPrice < float64(givenAmounth) {
		return false
	}
	return true
}
func (c *CartService) CalculateForthOrderDiscount(productList []product.Product) (float64, float64) {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	productCountById := make(map[product.Product]int)
	for _, product := range productList {
		productCountById[product]++
	}

	for selectedProduct := range productCountById {
		switch selectedProduct.ProductInfo.Vat {
		case 1:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			totalPrice += selectedProduct.ProductInfo.Price + vat
		case 8:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat - vat*0.1
			totalPrice += (selectedProduct.ProductInfo.Price + vat) - (selectedProduct.ProductInfo.Price+vat)*0.1
		case 18:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat - vat*0.15
			totalPrice += (selectedProduct.ProductInfo.Price + vat) - (selectedProduct.ProductInfo.Price+vat)*0.15
		}
	}
	return totalPrice, vatOfCart
}
