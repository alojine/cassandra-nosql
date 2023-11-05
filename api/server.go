package api

import (
	"encoding/json"
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
	// http.HandleFunc("/factory", s.handleGetFactoryById)
	http.HandleFunc("/products/{product_line_id}", s.handleGetProductsByProductLineId)
	return http.ListenAndServe(s.listenAddr, nil)
}

// func (s *Server) handleGetFactoryById(w http.ResponseWriter, r *http.Request) {
// 	factory := s.store.Get("1")
// 	json.NewEncoder(w).Encode(factory)
// }

func (s *Server) handleGetProductsByProductLineId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["product_line_id"]
	if !ok {
		http.Error(w, "product_line_id not found in request", http.StatusBadRequest)
		return
	}
	products := s.database.GetAllProductsByProductLineID(id)
	json.NewEncoder(w).Encode(products)
}
