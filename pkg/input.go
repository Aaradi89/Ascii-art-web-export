package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// Import the ascii banner to an arry of strings
func ReadArt(filename string) ([]string, bool) {

	var ascii []string
	valid := true
	switch filename {
	case "standard":
		filename = "Banners/standard.txt"
	case "shadow":
		filename = "Banners/shadow.txt"
	case "thinkertoy":
		filename = "Banners/thinkertoy.txt"
		// default:
		// 	valid = false
		// 	return ascii, valid
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		valid = false
		return ascii, valid
		//os.Exit(1)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		ascii = append(ascii, scan.Text())
	}
	if len(ascii) != 855 {
		valid = false
	}
	return ascii, valid
}

// Prepare the ascii banner to an arry of strings
func ArtPreparation(text string, art []string) []string {

	result := make([]string, 8)
	var prev rune

	for _, r := range text {
		if r >= ' ' && r <= '~' {
			position := ((int(r) - ' ') * 9) + 1

			for i := 0; i < 8; i++ {
				result[i] += art[position+i]
			}

		} else if r == 13 {
			// skip this & do noting

		} else if r == 10 && prev != 0 && prev == 13 { // We've a new line to print
			// skip this & do noting

		} else {
			// err := ArtPreparation("INPUT ERROR", art)
			err := []string{"Err", "Err", "Err", "Err", "Err", "Err", "Err", "Err"}
			return err

		}
		prev = r // store the previous rune for the next iteration
	}
	return result
}
