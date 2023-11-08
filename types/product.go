package types

type Product struct {
	ProductID        string `json:"product_id"`
	Name             string `json:"name"`
	Quantity int `json:quantity`
}

func (p Product) GetProductID() string {
	return p.ProductID
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) GetQuantity() int {
	return p.Quantity
}
