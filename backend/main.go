package main

import (
	"expense-tracker/database"
	"expense-tracker/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
)

func main() {
	log.Println("before db init")
	database.InitDB()
	log.Println("after db init")
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

	handler := c.Handler(http.DefaultServeMux)
	log.Println("before handler call")

	http.HandleFunc("/api/dashboard", handlers.DashboardHandler)
	http.HandleFunc("/api/expense", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.AddExpenseHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteExpenseHandler(w, r)
		case http.MethodPut:
			handlers.EditExpenseHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/auth/login", handlers.LoginHandler)
	http.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	http.ListenAndServe(":8080", handler)
}
