package api

import (
	"encoding/json"
	"example/NO-SQL-Cassandra/storage"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	database   storage.Database
}

func NewServer(listenAddr string, store storage.Storage, database storage.Database) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		database:   database,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/factory", s.handleGetFactoryById)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetFactoryById(w http.ResponseWriter, r *http.Request) {
	factory := s.store.Get("1")
	json.NewEncoder(w).Encode(factory)
}
