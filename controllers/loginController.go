package controllers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/both/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Title":   "Trang chu",
		"Heading": "Chao mung den voi IT Admin",
		"Message": "Day la he thong quan tri IT",
	}
	tmpl.Execute(w, data)
}
