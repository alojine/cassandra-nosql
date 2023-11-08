package types

type ProductLinesByFactory struct{
	FactoryName             string `json:"factory_name"`
	ProductLineName             string `json:"product_name"`
	ProductLineID       string `json:"product_line_id"`
}