package entity_cart

type CartInfo struct {
	TotalPrice float64 `json:"price"`
	VatOfCart  float64 `json:"vat"`
}

func (cartinfo *CartInfo) GetPrice() float64 {
	return cartinfo.TotalPrice
}

func (cartinfo *CartInfo) SetPrice(Price float64) *CartInfo {
	cartinfo.TotalPrice = Price
	return cartinfo
}

func (cartinfo *CartInfo) GetVat() float64 {
	return cartinfo.VatOfCart
}
func (cartinfo *CartInfo) SetVat(Vat float64) *CartInfo {
	cartinfo.VatOfCart = Vat
	return cartinfo
}
