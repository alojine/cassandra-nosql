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

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port", *listenAddr)
	log.Fatal(server.Start())
}
