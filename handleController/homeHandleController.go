package handleController

import (
	"html/template"
	"net/http"
)

func HomeHandleGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
