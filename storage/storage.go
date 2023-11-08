package storage

import "example/NO-SQL-Cassandra/types"

type Storage interface {
	Get(string) *types.Factory
	GetAllProductsByProductLineID(string) []*types.Product
}
