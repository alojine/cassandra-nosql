package types

type ProductionLine struct {
	LineID      string `json:"line_id"`
	FactoryID   string `json:"factory_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
