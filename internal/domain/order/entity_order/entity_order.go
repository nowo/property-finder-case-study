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
