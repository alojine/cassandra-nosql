package storage

import "example/NO-SQL-Cassandra/types"

type Storage interface {
	Get(string) *types.Factory
	// PProduct by product line
	// Product_line by factory
	GetAllProductsByProductLineID(string) []*types.Product
}
