// Package handlers contains the function for handling GET requests for the ASCII form submission.
package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// GetAbout handles GET requests for the "/about" page.
// It returns an error response for non-GET requests and other paths.

func GetAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderErrorPage(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Attempt to parse the template file.

	t, err := template.ParseFiles("templates/about.html")
	if err != nil {
		// Log parsing errors and return a 500 Internal Server Error.
		log.Printf("Error parsing template: %v\n", err)
		renderErrorPage(w, "404 Not Found", http.StatusNotFound)
		return
	}
	// Execute the parsed template and write it to the response.
	err = t.Execute(w, nil)
	if err != nil {
		// Log execution errors and return a 500 Internal Server Error.
		log.Printf("Error executing template: %v\n", err)
		renderErrorPage(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
