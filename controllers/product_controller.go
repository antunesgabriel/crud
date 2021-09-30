package controllers

import (
	"net/http"

	"github.com/antunesgabriel/crud/configs"
	"github.com/antunesgabriel/crud/models"
)

type ProductController struct {
}

var view = configs.View{}

func (*ProductController) Index(wr http.ResponseWriter, req *http.Request) {
	products, _ := new(models.Product).GetAll()

	view.Load(wr, "Index", products)
}
