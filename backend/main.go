package main

import (
	"expense-tracker/database"
	"expense-tracker/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	// Graceful shutdown handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Shutting down server...")
		database.CloseDB()
		os.Exit(0)
	}()

	log.Println("Starting server on port 8080...")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	mux := routes.RegisterRoutes()

	handler := c.Handler(mux)
	log.Println("before handler call")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err, " - Error starting server")
		return
	}
}
