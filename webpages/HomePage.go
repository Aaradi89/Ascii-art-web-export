package webpages

import (
	"html/template"
	"net/http"
)

func HomePage(ResWriter http.ResponseWriter, r *http.Request) { // Load the main/home page

	if r.Method != http.MethodGet {
		ResWriter.WriteHeader(400)
		t, _ := template.ParseFiles("templates/400.html") // If this page entered without input text
		t.Execute(ResWriter, nil)
	} else {

		switch r.URL.Path {
		case "/":
			result := data{Visibility: false}                      // hide the art box
			t, _ := template.ParseFiles("templates/HomePage.html") //To enter the home page
			t.Execute(ResWriter, result)
		default:
			ResWriter.WriteHeader(404)
			t, _ := template.ParseFiles("templates/404.html") // If user entered an invalid URL
			t.Execute(ResWriter, nil)
		}
	}
}
