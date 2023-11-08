package storage

import (
	"example/NO-SQL-Cassandra/types"
	"fmt"
)

func (db *Database) GetAllProductsByProductLineName(ProductLineName string) []*types.ProductsByProductLine {
	var products []*types.ProductsByProductLine

	iter := db.session.Query(`SELECT product_line_name, product_name, product_id FROM Products_By_Product_Line WHERE product_line_name = ?;`, ProductLineName).Iter()

	for {
		product := &types.ProductsByProductLine{}
		if !iter.Scan(&product.ProductLineName, &product.ProductName, &product.ProductID) {
			break
		}

		products = append(products, product)
	}

	return products
}

func (db *Database) GetAllProductLinesByFactoryName(FactoryName string) []*types.ProductLinesByFactory {
	var productLines []*types.ProductLinesByFactory

	iter := db.session.Query(`SELECT factory_name, product_line_name, product_line_id FROM Product_Lines_By_Factory WHERE factory_name = ?;`, FactoryName).Iter()

	for {
		productLine := &types.ProductLinesByFactory{}
		if !iter.Scan(&productLine.FactoryName, &productLine.ProductLineName, &productLine.ProductLineID) {
			break
		}

		productLines = append(productLines, productLine)
	}

	return productLines
}

func (db *Database) InsertProduct(ProductID string, Name string, Quantity string) bool {
	query := db.session.Query(
		`INSERT INTO Products (product_id, name, quantity) VALUES (?, ?, ?) IF NOT EXISTS;`,
		ProductID, Name, Quantity,
	)

	var applied bool

	if err := query.Scan(&applied); err != nil {
		fmt.Println("Product inserted successfully")
	} else {
		fmt.Println("Product with the same product_id already exists.")
		updateQuery := db.session.Query(
			"UPDATE Products SET quantity = ? IF name = ?;",
			Quantity, Name,
		)
		if err := updateQuery.Exec(); err != nil {
			fmt.Println("Product id matches, but product name does not match.")
			return false
		}

		return true
	}

	return true
}
