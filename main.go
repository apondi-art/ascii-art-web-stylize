package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ascii-art-web/handlers"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Usage: 'go run .'")
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.GetAsciiForm)
	http.HandleFunc("/about", handlers.GetAbout)
	http.HandleFunc("/ascii-art", handlers.PostAsciiArt)
	fmt.Println("SUCCESS!! listen to server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
