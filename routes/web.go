package routes

import (
	"net/http"

	"github.com/antunesgabriel/crud/controllers"
)

func LoadRoutes() {
	productController := new(controllers.ProductController)

	http.HandleFunc("/", productController.Index)
	http.HandleFunc("/new", productController.Create)
	http.HandleFunc("/store", productController.Store)
}
