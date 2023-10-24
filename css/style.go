package style

import "net/http"

func HomePageStyle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "css/HomePageStyle.css")
}

func PostPageStyle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "css/PostPageStyle.css")
}
