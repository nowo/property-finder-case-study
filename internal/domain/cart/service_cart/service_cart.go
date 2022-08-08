package service_cart

import (
	"fmt"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"property-finder-go-bootcamp-homework/internal/domain/cart/repository_cart"
	"property-finder-go-bootcamp-homework/internal/domain/order/repository_order"
	"property-finder-go-bootcamp-homework/internal/domain/product"
	"property-finder-go-bootcamp-homework/internal/domain/product/repository_product"
	"property-finder-go-bootcamp-homework/pkg/constants"
	_jwt "property-finder-go-bootcamp-homework/pkg/jwt"
)

type CartService struct {
	CartRepo    repository_cart.ICartRepository
	ProductRepo repository_product.IProductRepository
	OrderRepo   repository_order.IRepositoryOrder
	jwt         _jwt.JWT
}

func New(cartRepo repository_cart.ICartRepository, productRepo repository_product.IProductRepository, orderRepo repository_order.IRepositoryOrder) ICartService {
	return &CartService{
		CartRepo:    cartRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
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

func (c *CartService) CalculatePrice(cartList []product.Product, userID uint) (float64, float64) {

	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	for _, product := range cartList {
		vat := (product.ProductInfo.Price * float64(product.ProductInfo.Vat)) / 100
		vatOfCart += vat
		totalPrice += product.ProductInfo.Price + vat
	}

	discountAPrice, discountBPrice, discountCPrice := totalPrice, totalPrice, totalPrice
	discountAVat, discountBVat, discountCVat := vatOfCart, vatOfCart, vatOfCart

	discountAPrice, discountAVat = c.applyForthOrderDiscountPrice(cartList, userID, discountAPrice, discountAVat)
	discountBPrice, discountBVat = c.applySameProductDiscountPrice(cartList)
	discountCPrice, discountCVat = c.applyMonthlyDiscountPrice(userID, totalPrice, vatOfCart)

	return getSmallestDiscount(discountAPrice, discountBPrice, discountCPrice), getSmallestDiscount(discountAVat, discountBVat, discountCVat)
}

func getSmallestDiscount(discountA, discountB, discountC float64) float64 {
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

func (c *CartService) applySameProductDiscountPrice(productList []product.Product) (float64, float64) {
	var totalPrice float64 = 0
	var vatOfCart float64 = 0

	productCountById := make(map[product.Product]int)
	for _, product := range productList {
		productCountById[product]++
	}
	fmt.Println("productmap")
	fmt.Println(productCountById)
	fmt.Println("productlist")
	fmt.Println(productList)

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
