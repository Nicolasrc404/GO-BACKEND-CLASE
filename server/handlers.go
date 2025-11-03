package server

import (
	"backend-avanzado/api"
	"backend-avanzado/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllPeople(w, r)
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
		Name: p.Name,
		Age:  p.Age,
	}
	person, err = s.PeopleRepository.Save(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pResponse := &api.PersonResponse{
		ID:        int(person.ID),
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

func (s *Server) handleGetAllPeople(w http.ResponseWriter, r *http.Request) {
	response := make([]*api.PersonResponse, 0)
	people, err := s.PeopleRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, v := range people {
		personResponde := &api.PersonResponse{
			ID:        int(v.ID),
			Name:      v.Name,
			Age:       v.Age,
			CreatedAt: v.CreatedAt.String(),
		}
		response = append(response, personResponde)
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

	person, err := s.PeopleRepository.FindById(int(id))
	if person != nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &api.PersonResponse{
		ID:        int(person.ID),
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
