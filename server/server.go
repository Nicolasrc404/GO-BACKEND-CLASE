package server

import (
	"backend-avanzado/logger"
	"backend-avanzado/models"
	"fmt"
	"net/http"
)

type Server struct {
	DB     []*models.Person
	Logger *logger.Logger
}

func (s *Server) StartServer() {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: s.router(),
	}
	fmt.Println("Escuchando en el puerto 8000...")
	if err := srv.ListenAndServe(); err != nil {
		// Mostrar error
		s.Logger.Fatal(err)
	}
}

func NewServer() *Server {
	return &Server{
		Logger: logger.NewLogger(),
		DB:     make([]*models.Person, 0),
	}
}
