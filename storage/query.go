package storage

import (
	"example/NO-SQL-Cassandra/types"
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