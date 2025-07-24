package routes

import (
	"net/http"

	"github.com/vPt-King/ITMGMT/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/home", controllers.HomeHandler)
}
