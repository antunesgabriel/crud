package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (*ProductController) Create(wr http.ResponseWriter, req *http.Request) {
	view.Load(wr, "New", nil)
}

func (*ProductController) Store(wr http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.PostFormValue("name")
		description := req.PostFormValue("description")
		priceString := req.PostFormValue("price")
		priceString = strings.Replace(priceString, ",", ".", 1)
		price, err := strconv.ParseFloat(priceString, 64)

		if err != nil {
			fmt.Println("Error on convert price", err)
			http.Redirect(wr, req, "/", http.StatusBadRequest)
		}

		amount, err := strconv.Atoi(req.PostFormValue("amount"))

		if err != nil || name == "" || description == "" {
			fmt.Println("Invalid values", err)
			http.Redirect(wr, req, "/", http.StatusBadRequest)
		}

		err = new(models.Product).Store(&name, &description, &price, &amount)

		if err != nil {
			fmt.Println("Fail on insert product", err)
			http.Redirect(wr, req, "/", http.StatusBadRequest)
		}

	}

	http.Redirect(wr, req, "/", http.StatusMovedPermanently)
}
