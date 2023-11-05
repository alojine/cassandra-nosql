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
	r := mux.NewRouter()
	http.HandleFunc("/factory", s.handleGetFactoryById)
	http.HandleFunc("/products", s.handleGetProductsByProductLineId)

	http.Handle("/", r)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetFactoryById(w http.ResponseWriter, r *http.Request) {
	factory := s.database.Get("1")
	json.NewEncoder(w).Encode(factory)
}

func (s *Server) handleGetProductsByProductLineId(w http.ResponseWriter, r *http.Request) {
	productLineID := r.URL.Query().Get("product_line_id")
	if productLineID == "" {
        http.Error(w, "product_line_id not found in query parameters", http.StatusBadRequest)
        return
    }
	products := s.database.GetAllProductsByProductLineID(productLineID)
	json.NewEncoder(w).Encode(products)
}
