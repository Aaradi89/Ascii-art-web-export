package main

import (
	style "ascii-art-web/css"
	"ascii-art-web/webpages"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/img/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/500", webpages.Error500)
	http.HandleFunc("/", webpages.HomePage)                    // To handle the home/main page - request handler in / path
	http.HandleFunc("/ascii-art", webpages.PostRequest)        // To handle the POST request in /ascii-art path
	http.HandleFunc("/HomePageStyle.css", style.HomePageStyle) // To handle HomePage style in css
	http.HandleFunc("/PostPageStyle.css", style.PostPageStyle) // To handle PostPage Style in css

	go func() { //To run the server in the background
		log.Fatal(http.ListenAndServe(":8080", nil)) // Start the http server & listen on port number 8080 - If error, log will show it
		//? Nothing will be executed here
	}()

	fmt.Println("------------------------------------------------\nServer Starting...\nUse the following link : http://localhost:8080/\n-----------------------------------------------")

	select {} // Prevent the program from exit & keep the server running
}
