package api

import (
	"example/NO-SQL-Cassandra/storage"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	database   storage.Database
}

func NewServer(listenAddr string, database storage.Database) *Server {
	return &Server{
		listenAddr: listenAddr,
		database:   database,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	http.HandleFunc("/products", s.handleGetProductsByProductLineName)
	http.HandleFunc("/product-lines", s.handleGetProductLinesByFactoryName)
	http.Handle("/", r)

	return http.ListenAndServe(s.listenAddr, nil)
}
