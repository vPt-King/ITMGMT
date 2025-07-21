package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthMethod string `json:"auth_method"`
}

func LoginGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/both/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Template error:", err) // Thêm dòng này để debug

		return
	}
	data := map[string]interface{}{
		"Title":   "Trang chu",
		"Heading": "Chao mung den voi IT Admin",
		"Message": "Day la he thong quan tri IT",
	}
	tmpl.Execute(w, data)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	authMethod := r.FormValue("auth_method")

	user := User{Username: username, Password: password, AuthMethod: authMethod}

	var users []User
	fileName := "users.json"
	if _, err := os.Stat(fileName); err == nil {
		data, _ := ioutil.ReadFile(fileName)
		json.Unmarshal(data, &users)
	}

	users = append(users, user)
	data, _ := json.MarshalIndent(users, "", "  ")
	ioutil.WriteFile(fileName, data, 0644)

	tmpl, err := template.ParseFiles("views/both/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := map[string]interface{}{
		"Title":   "Trang chu",
		"Heading": "Chao mung den voi IT Admin",
		"Message": "Đăng nhập thành công!",
	}
	tmpl.Execute(w, result)
}
