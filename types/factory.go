package types

type Factory struct {
	FactoryID string `json:"factory_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Capacity  int    `json:"capacity"`
}
