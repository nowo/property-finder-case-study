package entity_cart

//CartInfo contains all the information about the cart except id
type CartInfo struct {
	UserID      uint `json:"user_id"`
	ProductID   uint `json:"product_id"`
	OrderID     uint `json:"order_id"`
	IsCompleted bool `json:"is_completed"`
}

//Create new cartInfo and return it
func NewCartInfo(userID uint, productID uint, orderID uint, isCompleted bool) *CartInfo {
	cartInfo := new(CartInfo)
	cartInfo.UserID = userID
	cartInfo.ProductID = productID
	cartInfo.OrderID = orderID
	cartInfo.IsCompleted = isCompleted
	return cartInfo
}
