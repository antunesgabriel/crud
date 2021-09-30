package configs

import (
	"html/template"
	"net/http"
)

var resources = template.Must(template.ParseGlob("templates/*.html"))

type View struct{}

func (*View) Load(wr http.ResponseWriter, name string, data interface{}) {
	resources.ExecuteTemplate(wr, name, data)
}
