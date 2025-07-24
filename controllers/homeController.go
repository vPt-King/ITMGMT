package controllers

import (
	"net/http"

	handleController "github.com/vPt-King/ITMGMT/handleController"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleController.HomeHandleGet(w, r)
	}
}
