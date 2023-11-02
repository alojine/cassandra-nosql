package storage

import "example/NO-SQL-Cassandra/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id string) *types.Factory {
	return &types.Factory{
		FactoryID: "1",
		Name:      "Foo",
		Location:  "LA",
		Capacity:  2000,
	}
}
