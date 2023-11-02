package storage

import "example/NO-SQL-Cassandra/types"

type CassandraStorage struct{}

func (s *CassandraStorage) Get(id string) *types.Factory {
	return &types.Factory{
		FactoryID: "1",
		Name:      "Foo",
		Location:  "LA",
		Capacity:  2000,
	}
}
