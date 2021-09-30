package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/antunesgabriel/crud/configs"
)

var resources = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":4200", nil); err != nil {
		log.Fatalln("Erro on stater server", err)
	}
}

func index(wr http.ResponseWriter, request *http.Request) {

	database := configs.ConnectDatabase()
	defer database.Close()

	result, err := database.Query("SELECT * from products")

	if err != nil {
		log.Fatalln("Error on query products", err)
	}

	product := Product{}
	products := []Product{}

	for result.Next() {
		var id, amount int
		var name, description string
		var price float64

		if err := result.Scan(&id, &name, &description, &price, &amount); err != nil {
			log.Fatalln("Error on scan row", err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	resources.ExecuteTemplate(wr, "Index", products)
}
