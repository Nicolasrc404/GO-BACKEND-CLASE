package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) defaultRoute(w http.ResponseWriter, r *http.Request) {
	type defaultResponse struct {
		Hola string `json:"hola"`
	}
	resp := &defaultResponse{
		Hola: "Mundo",
	}
	response, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(response)
}
