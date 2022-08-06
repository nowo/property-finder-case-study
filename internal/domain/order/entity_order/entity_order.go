package entity_order

type OrderInfo struct {
	UserID     uint    `json:"user_id"`
	TotalPrice float64 `json:"price"`
	VatOfCart  float64 `json:"vat"`
}

func NewOrderInfo(userID uint, totalPrice float64, vatOfCart float64) *OrderInfo {
	return &OrderInfo{
		UserID:     userID,
		TotalPrice: totalPrice,
		VatOfCart:  vatOfCart,
	}
}
func (orderinfo *OrderInfo) GetUserID() uint {
	return orderinfo.UserID
}

func (orderinfo *OrderInfo) GetTotalPrice() float64 {
	return orderinfo.TotalPrice
}

func (orderinfo *OrderInfo) GetVatOfCart() float64 {
	return orderinfo.VatOfCart
}

func (orderinfo *OrderInfo) SetUserID(UserID uint) {
	orderinfo.UserID = UserID
}

func (orderinfo *OrderInfo) SetTotalPrice(TotalPrice float64) {
	orderinfo.TotalPrice = TotalPrice
}

func (orderinfo *OrderInfo) SetVatOfCart(VatOfCart float64) {
	orderinfo.VatOfCart = VatOfCart
}
