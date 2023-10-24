package webpages

import (
	"ascii-art-web/pkg"
	"html/template"
	"net/http"
	"strings"
)

type data struct { // data pase to html template
	Art        string // the art to be printed
	Visibility bool   //art box status - false: hidden; or true: visible;
}

func PostRequest(w http.ResponseWriter, r *http.Request) { // Handle the POST request - Convert Text to ASCII ART

	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		t, _ := template.ParseFiles("templates/400.html") // If this page entered without input text
		t.Execute(w, nil)
		return
	}

	var combinedArt strings.Builder        // Store the final art  into one string
	inputText := r.FormValue("inputText")  // Store input text
	banner := r.FormValue("banner")        // Store the Selected Banner
	asciiArt, valid := pkg.ReadArt(banner) // Store the choosen banner content
	if !valid {
		w.WriteHeader(http.StatusInternalServerError)
		t, _ := template.ParseFiles("templates/500.html")
		t.Execute(w, nil) //Transfer to 500 Error Page
	} else {
		// tmpl, _ := template.ParseFiles("html/HomePage.html")
		testNotAllowedMethod := false

		if pkg.IsNewLine(inputText) { // If there is a new line in the input text

			Words := strings.Split(inputText, "\n") //Split the text if user entered a new lines in between

			for x := 0; x < len(Words); x++ {

				finalArt := pkg.ArtPreparation(Words[x], asciiArt) // Store the final ASCII art for this word
				if finalArt[0] == "Err" {
					testNotAllowedMethod = true
					break
				}
				for _, line := range finalArt {
					combinedArt.WriteString(line + "\n") // Store all string bytes into one string seperated by new line

				}

			}
			// resultString := combinedArt.String() // Convert the stored values into string
			// tmpl.Execute(w, resultString)        //Write the results into the webpage

		} else if inputText != "" { // If there is NO new line in the input text

			finalArt := pkg.ArtPreparation(inputText, asciiArt) // Store the final ASCII art for the Input text
			if finalArt[0] == "Err" {
				testNotAllowedMethod = true
			} else {
				for _, line := range finalArt {
					combinedArt.WriteString(line + "\n") // Store all string bytes into one string seperated by new line
				}

				// resultString := combinedArt.String() // Convert the stored values into string
				// tmpl.Execute(w, resultString)        //Write the results into the webpage
			}
		}
		if testNotAllowedMethod {
			w.WriteHeader(http.StatusInternalServerError)
			t, _ := template.ParseFiles("templates/500.html")
			t.Execute(w, nil) //Transfer to 500 Error Page
		} else {
			tmpl, _ := template.ParseFiles("templates/HomePage.html")
			resultString := combinedArt.String()                // Convert the stored values into string
			result := data{Art: resultString, Visibility: true} // export art and change art box visibility to visible
			tmpl.Execute(w, result)                             //Write the results into the webpage
		}
	}
}
