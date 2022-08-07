package entity_cart

type CartInfo struct {
	UserID      uint `json:"user_id"`
	ProductID   uint `json:"product_id"`
	OrderID     uint `json:"order_id"`
	IsCompleted bool `json:"is_completed"`
}

func NewCartInfo(userID uint, productID uint, orderID uint, isCompleted bool) *CartInfo {
	cartInfo := new(CartInfo)
	cartInfo.SetUserID(userID)
	cartInfo.SetProductID(productID)
	cartInfo.SetOrderID(orderID)
	cartInfo.SetIsCompleted(isCompleted)
	return cartInfo
}
func (cartinfo *CartInfo) GetUserID() uint {
	return cartinfo.UserID
}

func (cartinfo *CartInfo) GetProductID() uint {
	return cartinfo.ProductID
}

func (cartinfo *CartInfo) GetOrderID() uint {
	return cartinfo.OrderID
}

func (cartinfo *CartInfo) GetIsCompleted() bool {
	return cartinfo.IsCompleted
}

func (cartinfo *CartInfo) SetUserID(UserID uint) {
	cartinfo.UserID = UserID
}

func (cartinfo *CartInfo) SetProductID(ProductID uint) {
	cartinfo.ProductID = ProductID
}

func (cartinfo *CartInfo) SetOrderID(OrderID uint) {
	cartinfo.OrderID = OrderID
}

func (cartinfo *CartInfo) SetIsCompleted(IsCompleted bool) {
	cartinfo.IsCompleted = IsCompleted
}
