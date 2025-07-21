package main

import (
	"fmt"
	"net/http"

	routes "github.com/vPt-King/ITMGMT/routers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	routes.RegisterRoutes()
	fmt.Println("Sever listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
