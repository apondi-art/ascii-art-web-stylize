// Package handlers contains the function for handling POST requests to generate ASCII art.
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	art "ascii-art-web/ascii-art"
)

// Data represents the data structure used for rendering the form template.
type Data struct {
	Filename string
	Input    string
	Result   string
}

// PostAsciiArt handles POST requests to the "/ascii-art" route.
// It processes the form submission to generate ASCII art or reset the form.
func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request on /ascii-art route\n", r.Method)

	if r.Method != http.MethodPost {
		renderErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Retrieve input values from the form submission.
	text := r.FormValue("input")
	banner := r.FormValue("filename")

	if text == "" || banner == "" {
		renderErrorPage(w, "Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}
	// Generate ASCII art using the provided input and banner filename.
	result, err := art.AsciiArt(text, banner)
	if err != nil {
		if strings.Contains(err.Error(), "Not within the printable ascii range") {
			log.Printf("Error generating ASCII art: %v\n", err)
			renderErrorPage(w, "Bad Request\n"+err.Error(), http.StatusBadRequest)
			return
		} else if strings.Contains(err.Error(), "no such file or directory") {
			log.Printf("Error generating ASCII art: %v\n", err)
			renderErrorPage(w, "No such file or directory:Not Found\n", http.StatusNotFound)
			return
		} else {
			log.Printf("Error generating ASCII art: %v\n", err)
			renderErrorPage(w, "Internal Server Error\n", http.StatusInternalServerError)
			return
		}
	}
	// Prepare the data for rendering the template with results.
	resultData := &Data{
		Filename: banner,
		Input:    text,
		Result:   result,
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		renderErrorPage(w, "Not Found : Missing Template file", http.StatusNotFound)
		return
	}

	err = t.Execute(w, resultData)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		renderErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
