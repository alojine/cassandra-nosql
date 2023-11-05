package storage

import (
	"example/NO-SQL-Cassandra/types"
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type Database struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func SetupDBConnection() (*Database, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "factory"
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to Cassandra on port :", cluster.Port)

	return &Database{
		cluster: cluster,
		session: session,
	}, nil
}

func (db *Database) Get(id string) *types.Factory {
	return &types.Factory{
		FactoryID: "1",
		Name:      "Foo",
		Location:  "LA",
		Capacity:  2000,
	}
}

func (db *Database) GetAllProductsByProductLineID(id string) []*types.Product {
	var products []*types.Product

	iter := db.session.Query(`SELECT product_id, product_line_id, name, description, production_date, status FROM products WHERE production_line_id = ?`, id).Iter()
	for {
		product := &types.Product{}
		if !iter.Scan(&product.ProductID, &product.ProductionLineID, &product.Name, &product.Description, &product.ProductionDate, &product.Status) {
			break
		}
		products = append(products, product)
		fmt.Println(product)
	}

	return products
}
