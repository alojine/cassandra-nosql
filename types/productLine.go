package types

type ProductLine struct {
	ProductLineID      int `json:"product_line_id"`
	ProductNames []string `json:"product_names"`
	Name        string `json:"name"`
}
