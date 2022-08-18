package service_cart

import (
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/pkg/constants"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
)

// ICartService interface contains all methods that are required to implement by service_cart.
type ICartService interface {
	AddToCart(userID, productID uint) error
	DeleteFromCart(userID, productID uint) error
	GetCartByUserID(userID uint) ([]product.Product, error)
	CalculatePrice(cartList []product.Product, userID uint) (float64, float64)
}

//cart service struct include cart repository, product repository, order repository and jwt
type CartService struct {
	CartRepo    repository_cart.ICartRepository
	ProductRepo repository_product.IProductRepository
	OrderRepo   repository_order.IOrderRepository
	jwt         _jwt.JWT
}

//Create new instance of cart service interface
func New(cartRepo repository_cart.ICartRepository, productRepo repository_product.IProductRepository, orderRepo repository_order.IOrderRepository) ICartService {
	return &CartService{
		CartRepo:    cartRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
		jwt:         *_jwt.New(),
	}
}

// AddToCart add product to cart if product quantity is not 0
func (c *CartService) AddToCart(userID, productID uint) error {

	newCart := cart.NewCart(userID, productID)

	selectedProduct, err := c.ProductRepo.GetProductByID(productID)
	if err != nil {
		return err
	}
	if selectedProduct.ProductInfo.Quantity == 0 {
		return messages.NOT_ENOUGH_QUANTITY
	}
	err = c.ProductRepo.UpdateProductQuantity(productID, selectedProduct.ProductInfo.Quantity-1)

	if err != nil {
		return err
	}
	if err := c.CartRepo.Create(*newCart); err != nil {
		return err
	}

	return nil
}

//Delete cart from product by userID and productID
func (c *CartService) DeleteFromCart(userID, productID uint) error {
	selectedProduct, err := c.ProductRepo.GetProductByID(productID)
	if err != nil {
		return err
	}

	err = c.ProductRepo.UpdateProductQuantity(productID, selectedProduct.ProductInfo.Quantity+1)
	if err != nil {
		return err
	}
	return c.CartRepo.Delete(userID, productID)
}

// GetCartByUserID return cart list by userID
func (c *CartService) GetCartByUserID(userID uint) ([]product.Product, error) {
	var buyedProducts = make([]product.Product, 0)
	cartList, err := c.CartRepo.GetCartsByUserID(userID)
	if err != nil {
		return nil, err
	}
	for _, cart := range cartList {
		selectedProduct, err := c.ProductRepo.GetProductByID(cart.CartInfo.ProductID)
		if err != nil {
			return nil, err
		}
		buyedProducts = append(buyedProducts, selectedProduct)
	}
	return buyedProducts, nil
}

//CalculateCartPrice calculate cart price and vat by comparing discount prices
func (c *CartService) CalculatePrice(cartList []product.Product, userID uint) (float64, float64) {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0
	for _, product := range cartList {
		vat := (product.ProductInfo.Price * float64(product.ProductInfo.Vat)) / 100
		vatOfCart += vat
		totalPrice += product.ProductInfo.Price + vat
	}

	forthOrderDiscountPrice, sameProductDiscountPrice, monthlyDiscountPrice := totalPrice, totalPrice, totalPrice
	forthOrderDiscountVat, sameProductDiscountVat, monthlyDiscountVat := vatOfCart, vatOfCart, vatOfCart

	forthOrderDiscountPrice, forthOrderDiscountVat = c.applyForthOrderDiscountPrice(cartList, userID, forthOrderDiscountPrice, forthOrderDiscountVat)
	sameProductDiscountPrice, sameProductDiscountVat = c.applySameProductDiscountPrice(cartList)
	monthlyDiscountPrice, monthlyDiscountVat = c.applyMonthlyDiscountPrice(userID, totalPrice, vatOfCart)

	return getSmallestTotalPriceAfterDiscount(forthOrderDiscountPrice, sameProductDiscountPrice, monthlyDiscountPrice), getSmallestTotalPriceAfterDiscount(forthOrderDiscountVat, sameProductDiscountVat, monthlyDiscountVat)
}

//compare discount prices and return smallest one
func getSmallestTotalPriceAfterDiscount(discountA, discountB, discountC float64) float64 {
	if discountA < discountB {
		if discountA < discountC {
			return discountA
		}
		return discountC
	}
	if discountB < discountC {
		return discountB
	}
	return discountC
}

//apply same product discount price if there is more than 3 same product, apply discount after third ones
func (c *CartService) applySameProductDiscountPrice(productList []product.Product) (float64, float64) {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	productCountById := make(map[product.Product]int)
	for _, product := range productList {
		productCountById[product]++
	}
	discountOrderCount := constants.SameProductDiscountCount - 1
	for selectedProduct, productCount := range productCountById {
		if productCount > discountOrderCount {
			//Todo: bu magic number olan 3u sil
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat*float64(discountOrderCount) + vat*float64(productCount-discountOrderCount)*0.8
			price := selectedProduct.ProductInfo.Price + vat
			totalPrice += price*float64(discountOrderCount) + price*float64(productCount-discountOrderCount)*0.8
		} else {
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			totalPrice += (selectedProduct.ProductInfo.Price + vat) * float64(productCount)
		}
	}
	return totalPrice, vatOfCart
}

//if user exceed given amount in a month, apply %10 discount for all subsequent ones
func (c *CartService) applyMonthlyDiscountPrice(userID uint, totalPrice float64, vatOfCart float64) (float64, float64) {
	orders, err := c.OrderRepo.GetOrderFromLastMonth(userID)
	if err != nil {
		return totalPrice, vatOfCart
	}
	var totalPricesFromLastMonth float64 = 0
	for _, order := range orders {
		totalPricesFromLastMonth += order.OrderInfo.TotalPrice
	}

	if totalPricesFromLastMonth < constants.GivenAmounth {
		return totalPrice, vatOfCart
	}
	return totalPrice - (totalPrice * 0.1), vatOfCart - (vatOfCart * 0.1)
}

//check is user is on forth order which is exceed given amount
func (c *CartService) isUserDeservedForthOrderDiscount(userID uint, totalPrice float64) bool {
	orders, err := c.OrderRepo.GetOrderByUserID(userID)
	if err != nil {
		return false
	}
	orderCount := 0
	for _, order := range orders {
		if order.OrderInfo.TotalPrice > constants.GivenAmounth {
			orderCount++
		}
	}
	if orderCount%constants.SameProductDiscountCount != constants.SameProductDiscountCount-1 {
		return false
	}
	if totalPrice < constants.GivenAmounth {
		return false
	}
	return true
}

//apply forth order discount price if user is on forth order and exceed given amount
func (c *CartService) applyForthOrderDiscountPrice(productList []product.Product, userID uint, totalPrice, vatOfCard float64) (float64, float64) {
	if !c.isUserDeservedForthOrderDiscount(userID, totalPrice) {
		return totalPrice, vatOfCard
	}
	var appliedTotalPrice float64 = 0
	var vatOfCart float64 = 0

	for _, selectedProduct := range productList {
		switch selectedProduct.ProductInfo.Vat {
		case 1:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat
			appliedTotalPrice += selectedProduct.ProductInfo.Price + vat
		case 8:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat - vat*0.1
			appliedTotalPrice += (selectedProduct.ProductInfo.Price + vat) - (selectedProduct.ProductInfo.Price+vat)*0.1
		case 18:
			vat := (selectedProduct.ProductInfo.Price * float64(selectedProduct.ProductInfo.Vat)) / 100
			vatOfCart += vat - vat*0.15
			appliedTotalPrice += (selectedProduct.ProductInfo.Price + vat) - (selectedProduct.ProductInfo.Price+vat)*0.15
		}
	}
	return appliedTotalPrice, vatOfCart
}
