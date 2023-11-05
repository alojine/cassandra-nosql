package storage

import (
	"log"
	"fmt"
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

func (db *Database) GetAllProductsByProductLineID(string id) []*types.Product {
	var products []types.Product
	var product types.Product

	iter := db.session.Query(`SELECT product_id, product_line_id, name, description, production_date, status FROM products WHERE product_line_id = ?`, id).Iter()
	for iter.Scan(&product.ProductID, &product.ProductionLineID, &product.Name, &product.ProductionDate, &product.Status) {
		products = append(products, product)
		fmt.Println(product)
	}

	return &products
}
