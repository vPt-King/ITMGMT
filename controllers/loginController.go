package controllers

import (
	"net/http"

	handleController "github.com/vPt-King/ITMGMT/handleController"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleController.LoginHandleGet(w, r)
	case http.MethodPost:
		handleController.LoginHandlePost(w, r)
	}
}
