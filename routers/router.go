package routes

import (
	"net/http"
	"project/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/secure/login", controllers.LoginHandler)
}
