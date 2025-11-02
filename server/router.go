package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/people", s.HandlePeople).Methods(http.MethodGet, http.MethodPost)
	return router
}
