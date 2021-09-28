package main

import (
	"html/template"
	"log"
	"net/http"
)

var resources = template.Must(template.ParseGlob("templates/*.html"))

type Products struct {
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
	products := []Products{
		{"SSH", "SSH 480gb", 239.65, 10},
		{"Samsung Smart TV 55", "UHD 4K 55AU7700, Processador Crystal 4K", 2999.00, 67},
	}
	resources.ExecuteTemplate(wr, "Index", products)
}
