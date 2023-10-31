package main

import (
	"example/NO-SQL-Cassandra/api"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("server running on port: ", *listenAddr)
	log.Fatal(server.Start())
}
