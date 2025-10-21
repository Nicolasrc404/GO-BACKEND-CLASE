package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type Server struct{}

func StartServer() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	stopper := make(chan struct{})
	go func() {
		<-done
		close(stopper)
	}()
	server, err := newServer()
	if err != nil {
		return err
	}
	return server.Start(stopper)
}

func newServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: s.router(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-stop
	return srv.Shutdown(context.Background())
}

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", s.defaultRoute)
	return router
}
