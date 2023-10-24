package webpages

import (
	"html/template"
	"net/http"
)

func Error500(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/500.html")
	w.WriteHeader(500)
	t.Execute(w, nil)
}
