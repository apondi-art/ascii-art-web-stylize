package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// ErrorData holds the data to be used in the error page template.
type ErrorData struct {
	StatusCode   int
	ErrorMessage string
}

// renderErrorPage renders an error page with a specific status code and message.
func renderErrorPage(w http.ResponseWriter, errorMessage string, statusCode int) {
	// Set the HTTP status code for the response.
	w.WriteHeader(statusCode)

	// Create an ErrorData instance to pass to the template.
	errorData := ErrorData{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}

	// Parse the error template file.
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Log and return an internal server error if template parsing fails.
		log.Printf("Error parsing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the error data.
	err = t.Execute(w, errorData)
	if err != nil {
		// Log and return an internal server error if template execution fails.
		log.Printf("Error executing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
