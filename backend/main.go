package main

import (
	"expense-tracker/handlers"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	print("Starting server on port 8080...\n")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(http.DefaultServeMux)
	http.HandleFunc("/api/dashboard", handlers.DashboardHandler)
	http.ListenAndServe(":8080", handler)
}
