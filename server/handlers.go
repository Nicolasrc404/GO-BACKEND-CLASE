package server

import (
	"backend-avanzado/api"
	"backend-avanzado/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func (s *Server) HandlePeopleWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.HandleGetPeopleById(w, r)
		return
		/* case http.MethodPut:
			s.handeEditPerson(w, r)
			return
		case http.MethodDelete:
			s.handleDeletePerson(w, r)
			return */
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
		ID:        len(s.DB) + 1,
		Name:      p.Name,
		Age:       p.Age,
		CreatedAt: time.Now(),
	}
	s.DB = append(s.DB, person)
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
	for _, person := range s.DB {
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

func (s *Server) HandleGetPeopleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if int(id) >= len(s.DB)-1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	person := s.DB[id-1]
	resp := &api.PersonResponse{
		ID:        person.ID,
		Name:      person.Name,
		Age:       person.Age,
		CreatedAt: person.CreatedAt.String(),
	}
	result, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
