package storage

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var schemaInitialized bool
var dataInitialized bool

func (db *Database) InitializeSchema() error {
	if schemaInitialized {
		log.Println("Schema has beem intialized")
		return nil
	}

	statements := []string{
		"CREATE TABLE IF NOT EXISTS factories (factory_id UUID PRIMARY KEY, name TEXT, location TEXT, capacity INT)",
		"CREATE TABLE IF NOT EXISTS product_lines (line_id UUID PRIMARY KEY, factory_id UUID, name TEXT, description TEXT)",
		"CREATE TABLE IF NOT EXISTS products (product UUID PRIMARY KEY, production_line_id UUID, name TEXT, description TEXT, production_date DATE, status TEXT)",
	}

	for _, statement := range statements {
		if err := db.session.Query(statement).Exec(); err != nil {
			return err
		}
	}

	schemaInitialized = true
	return nil
}

func (db *Database) InitializeData() error {
	if dataInitialized {
		log.Println("Data has been initialized.")
		return nil
	}

	if err := db.session.Query("INSERT INTO factories (factory_id, name, location, capacity) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), "Factory 1", "Location 1", 1000).Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO product_lines (line_id, factory_id, name, description) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), gocql.TimeUUID(), "Line 1", "Description 1").Exec(); err != nil {
		return err
	}

	if err := db.session.Query("INSERT INTO products (product, production_line_id, name, description, production_date, status) VALUES (?, ?, ?, ?, ?, ?)",
		gocql.TimeUUID(), gocql.TimeUUID(), "Product 1", "Product Description 1", time.Now(), "Available").Exec(); err != nil {
		return err
	}

	dataInitialized = true
	return nil
}
