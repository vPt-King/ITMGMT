package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	routes.RegisterRoutes()
	http.HandleFunc("/secure/login", loginHandler)
	fmt.Println("Sever listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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
