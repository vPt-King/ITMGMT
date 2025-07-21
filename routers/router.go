package routes

import (
	"net/http"

	"github.com/vPt-King/ITMGMT/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/secure/login/get", controllers.LoginGet)
	http.HandleFunc("/secure/login/post", controllers.LoginPost)

}
