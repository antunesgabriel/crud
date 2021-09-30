package routes

import (
	"net/http"

	"github.com/antunesgabriel/crud/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", new(controllers.ProductController).Index)
}
