package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	StatusCode   int
	ErrorMessage string
}

func renderErrorPage(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.WriteHeader(statusCode)
	errorData := ErrorData{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Printf("Error parsing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, errorData)
	if err != nil {
		log.Printf("Error executing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
