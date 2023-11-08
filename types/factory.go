package types

type Factory struct {
	FactoryID int `json:"factory_id"`
	ProductLines []string `json:"product_line_names"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
}
