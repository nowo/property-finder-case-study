package general

import "property-finder-go-bootcamp-homework/internal/domain/product"

type Token struct {
	Token string `json:"token"`
}

//General response structure
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//Order response data structure
type BasketResponse struct {
	Cart       []product.Product `json:"cart"`
	TotalPrice float64           `json:"price"`
	VatOfCart  float64           `json:"vat"`
}
