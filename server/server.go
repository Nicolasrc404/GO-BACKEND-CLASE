package server

import (
	"backend-avanzado/logger"
	"backend-avanzado/models"
	"backend-avanzado/repository"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB               *gorm.DB
	Logger           *logger.Logger
	PeopleRepository repository.Repository[models.Person]
}

func (s *Server) StartServer() {
	s.InitDB()
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

func (s *Server) InitDB() {
	dsm := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dsm), &gorm.Config{})
	if err != nil {
		s.Logger.Fatal(err)
	}
	s.DB = db
	s.DB.AutoMigrate(&models.Person{})
	s.PeopleRepository = repository.NewPeopleRepository(s.DB)
}

func NewServer() *Server {
	return &Server{
		Logger: logger.NewLogger(),
	}
}
