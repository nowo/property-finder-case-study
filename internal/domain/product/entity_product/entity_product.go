package entity_product

//ProductInfo struct include productID, productName, productPrice and productDescription. This struct is used to store product information.
type ProductInfo struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Vat         float64 `json:"vat"`
	Quantity    int     `json:"quantity"`
}
