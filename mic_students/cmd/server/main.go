package main

import (
	"log"

	"github.com/vitorsiqueirarecife/students/internal/server"
)

func main() {
	router := server.SetupRouter()

	port := ":8085"

	if err := router.Run(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
