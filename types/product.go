package types

type Product struct {
	ProductID        string `json:"product_id"`
	ProductionLineID string `json:"production_line_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProductionDate   string `json:"production_date"`
	Status           string `json:"status"`
}
