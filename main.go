package main

import (
	"example/NO-SQL-Cassandra/api"
	"example/NO-SQL-Cassandra/storage"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server")
	flag.Parse()

	dbconnection, err := storage.SetupDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra db: %V", err)
	}
	dbconnection.InitializeSchema()
	// dbconnection.InitializeData()
	database := dbconnection

	server := api.NewServer(*listenAddr, *database)
	fmt.Println("server running on port", *listenAddr)
	log.Fatal(server.Start())
}
