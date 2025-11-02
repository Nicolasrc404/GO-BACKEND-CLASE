package server

import (
	"backend-avanzado/api"
	"backend-avanzado/models"
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAtPeople(w, r)
		return
	case http.MethodPost:
		s.handeleCreatePerson(w, r)
		return
	}
}

func (s *Server) handeleCreatePerson(w http.ResponseWriter, r *http.Request) {
	var p api.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	person := &models.Person{
		ID:        len(s.BD) + 1,
		Name:      p.Name,
		Age:       p.Age,
		CreatedAt: time.Now(),
	}
	s.BD = append(s.BD, person)
	pResponse := &api.PersonResponse{
		ID:        person.ID,
		Name:      person.Name,
		Age:       person.Age,
		CreatedAt: person.CreatedAt.String(),
	}
	result, err := json.Marshal(pResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (s *Server) handleGetAtPeople(w http.ResponseWriter, r *http.Request) {
	response := make([]*api.PersonResponse, 0)
	for _, person := range s.BD {
		personResponse := &api.PersonResponse{
			ID:        person.ID,
			Name:      person.Name,
			Age:       person.Age,
			CreatedAt: person.CreatedAt.String(),
		}
		response = append(response, personResponse)
	}
	result, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
