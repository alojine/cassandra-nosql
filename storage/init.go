package storage

import (
	"fmt"

	"github.com/gocql/gocql"
)

var dataInitialized bool

func (db *Database) InitializeSchema() error {
	statements := []string{
		"CREATE TABLE IF NOT EXISTS Factories(factory_id UUID, product_line_names SET<TEXT>, name TEXT, capacity INT, PRIMARY KEY(factory_id));",
		"CREATE TABLE IF NOT EXISTS Product_Lines(product_line_id UUID, product_names SET<TEXT>, name TEXT, PRIMARY KEY(product_line_id));",
		"CREATE TABLE IF NOT EXISTS Products(product_id UUID, name TEXT, quantity INT, PRIMARY KEY(product_id));",
		"CREATE TABLE IF NOT EXISTS Products_By_Product_Line (product_line_name TEXT, product_name TEXT, product_id UUID, PRIMARY KEY((product_line_name), product_name));",
		"CREATE TABLE IF NOT EXISTS Product_Lines_By_Factory (factory_name TEXT, product_line_name TEXT, product_line_id UUID, PRIMARY KEY((factory_name), product_line_name));",
	}

	for _, statement := range statements {
		if err := db.session.Query(statement).Exec(); err != nil {
			return err
		}
	}

	fmt.Println("Schema initialization has finished.")
	return nil
}

func (db *Database) InitializeData() error {

	factoryID, err := gocql.ParseUUID("743c227f-3903-4103-8bcb-09f76c38dc08")
	if err != nil {
		return err
	}
	productLineID, err := gocql.ParseUUID("c59d1ef6-70d2-4abd-9b80-5371b3314a64")
	if err != nil {
		return err
	}
	productID, err := gocql.ParseUUID("ecc6cce2-2c0d-4567-9bb5-e9cffd52ee64")
	if err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO factories (factory_id, product_line_names, name, capacity) VALUES (?, ?, ?, ?)",
		factoryID, []string{"Technology"}, "Apple", 1000).Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO product_lines (product_line_id, product_names, name) VALUES (?, ?, ?)",
		productLineID, []string{"IphoneX"}, "Technology").Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO products (product_id, name, quantity) VALUES (?, ?, ?)",
		productID, "IphoneX", 500).Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO products_by_product_line (product_line_name, product_name, product_id) VALUES (?, ?, ?)",
		"Technology", "IphoneX", productID).Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO product_lines_by_factory (factory_name, product_line_name, product_line_id) VALUES (?, ?, ?)",
		"Apple", "Technology", productLineID).Exec(); err != nil {
		return err
	}

	fmt.Println("Data initialization has finished.")
	return nil
}
