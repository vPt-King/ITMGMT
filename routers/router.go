package routes

import (
	"net/http"

	"github.com/vPt-King/ITMGMT/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/secure/login", controllers.LoginHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
}
