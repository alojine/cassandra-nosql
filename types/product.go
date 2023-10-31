package types

type Product struct {
	ProductID        string `json:"product_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProductionLineID string `json:"production_line_id"`
	ProductionDate   string `json:"production_date"`
	Status           string `json:"status"`
}
