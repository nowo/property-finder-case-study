package entity

type ProductInfo struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Vat         float64 `json:"vat"`
	Quantity    int     `json:"quantity"`
}
