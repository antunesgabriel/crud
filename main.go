package main

import (
	"log"
	"net/http"

	"github.com/antunesgabriel/crud/routes"
)

func main() {
	routes.LoadRoutes()

	if err := http.ListenAndServe(":4200", nil); err != nil {
		log.Fatalln("Erro on stater server", err)
	}
}
