package main

import (
	"backend-avanzado/server"
	"log"
	"os"
)

func main() {
	if err := server.StartServer(); err != nil {
		log.Fatal("Error iniciando el servidor: %v", err)
		os.Exit(1)
	}
}
