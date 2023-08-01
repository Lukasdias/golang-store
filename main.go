package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name string
	Price float64
	Quantity int
	Description string
}

var temp = template.Must(template.ParseGlob("templates/*.html"))


func main() {
	http.HandleFunc("/",  index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Product{
		{"Mouse", 25.00, 10, "Mouse sem fio"},
		{"Teclado", 50.00, 10, "Teclado sem fio"},
		{"Monitor", 500.00, 10, "Monitor 4k"},
		{"Notebook", 2500.00, 10, "Notebook Dell"},
		{"Impressora", 500.00, 10, "Impressora HP"},
		{"Cadeira", 250.00, 10, "Cadeira Gamer"},
		{"Mesa", 500.00, 10, "Mesa Gamer"},
		{"Gabinete", 250.00, 10, "Gabinete Gamer"},
		{"Placa de Vídeo", 2500.00, 10, "Placa de Vídeo RTX 2080"},
	}
	temp.ExecuteTemplate(w, "index.html", produtos)
}