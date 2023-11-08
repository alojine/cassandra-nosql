package api

import (
	"net/http"
	"encoding/json"
)

func (s *Server) handleGetProductLinesByFactoryName(w http.ResponseWriter, r *http.Request) {
	factoryName := r.URL.Query().Get("factory_name")

	if factoryName == "" {
        http.Error(w, "factory_name not found in query parameters", http.StatusBadRequest)
        return
    }
	productLines := s.database.GetAllProductLinesByFactoryName(factoryName)
	json.NewEncoder(w).Encode(productLines)
}

func (s *Server) handleGetProductsByProductLineName(w http.ResponseWriter, r *http.Request) {
	productLineName := r.URL.Query().Get("product_line_name")

	if productLineName == "" {
        http.Error(w, "product_line_name not found in query parameters", http.StatusBadRequest)
        return
    }
	products := s.database.GetAllProductsByProductLineName(productLineName)
	json.NewEncoder(w).Encode(products)
}