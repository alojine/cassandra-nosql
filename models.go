package models

type Factory struct {
	FactoryID string `json:"factory_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Capacity  int    `json:"capacity"`
}

type ProductionLine struct {
	LineID      string `json:"line_id"`
	FactoryID   string `json:"factory_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	ProductID        string `json:"product_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProductionLineID string `json:"production_line_id"`
	ProductionDate   string `json:"production_date"`
	Status           string `json:"status"`
}
