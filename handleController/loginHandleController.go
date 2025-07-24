package handleController

import (
	"html/template"
	"net/http"
)

func LoginHandleGet(w http.ResponseWriter, r *http.Request) {
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

func LoginHandlePost(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	// password := r.FormValue("password")
	// authMethod := r.FormValue("authMethod")
	// if authMethod == "local" {
	// 	if username == "admin" && password == "admin" {
	// 		session := model.Session{utils.GenerateSessionID(), "123456", "admin"}
	// 		cookie := http.Cookie{
	// 			Name:     "session_id",
	// 			Value:    session.Id,
	// 			Path:     "/",
	// 			HttpOnly: true,
	// 			Secure:   false,
	// 			SameSite: http.SameSiteDefaultMode,
	// 			Expires:  time.Now().Add(24 * time.Hour),
	// 		}

	// 	}
	// 	var sessionId, err = utils.GenerateSessionID()
	// 	if err != nil {
	// 		http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	tmpl, err := template.ParseFiles("views/login.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	data := map[string]interface{}{
	// 		"Title":   "Trang chu",
	// 		"Heading": "Chao mung den voi IT Admin",
	// 		"Message": "Đăng nhập thành công!",
	// 	}
	// 	tmpl.Execute(w, data)
	// 	return
	// }
}
