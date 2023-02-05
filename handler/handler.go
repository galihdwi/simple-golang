package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// 404 handler
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Golang Web",
	// 	"content": "Hello World",
	// 	"version": 1.0,
	// }

	data := []entity.Product{
		{ID: 1, Name: "Apel", Price: 1000, Stock: 10},
		{ID: 2, Name: "Mangga", Price: 2000, Stock: 20},
		{ID: 3, Name: "Semangka", Price: 3000, Stock: 30},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error on web", http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	id_number, erro := strconv.Atoi(id)

	if erro != nil || id_number < 1 {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"id": id_number,
	}

	view, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = view.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error on web", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// log.Println("Product Page" + id)
}
