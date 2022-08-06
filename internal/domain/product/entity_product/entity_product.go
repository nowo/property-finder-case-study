package entity_product

type ProductInfo struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Vat         float64 `json:"vat"`
	Quantity    int     `json:"quantity"`
}

func (productinfo *ProductInfo) GetName() string {
	return productinfo.Name
}

func (productinfo *ProductInfo) GetPrice() float64 {
	return productinfo.Price
}

func (productinfo *ProductInfo) GetDescription() string {
	return productinfo.Description
}

func (productinfo *ProductInfo) GetVat() float64 {
	return productinfo.Vat
}

func (productinfo *ProductInfo) GetQuantity() int {
	return productinfo.Quantity
}

func (productinfo *ProductInfo) SetName(Name string) *ProductInfo {
	productinfo.Name = Name
	return productinfo
}

func (productinfo *ProductInfo) SetPrice(Price float64) *ProductInfo {
	productinfo.Price = Price
	return productinfo
}

func (productinfo *ProductInfo) SetDescription(Description string) *ProductInfo {
	productinfo.Description = Description
	return productinfo
}

func (productinfo *ProductInfo) SetVat(Vat float64) *ProductInfo {
	productinfo.Vat = Vat
	return productinfo
}

func (productinfo *ProductInfo) SetQuantity(Quantity int) *ProductInfo {
	productinfo.Quantity = Quantity
	return productinfo
}
