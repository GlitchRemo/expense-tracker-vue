package routes

import (
	"expense-tracker/handlers"
	"net/http"
)

func RegisterRoutes() http.Handler {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/api/dashboard", handlers.DashboardHandler)
	serveMux.HandleFunc("/api/expense", func(w http.ResponseWriter, r *http.Request) {
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
	serveMux.HandleFunc("/api/auth/login", handlers.LoginHandler)
	serveMux.HandleFunc("/api/auth/register", handlers.RegisterHandler)

	// Category routes
	serveMux.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetCategories(w, r)
		case http.MethodPost:
			handlers.CreateCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	serveMux.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		// expects /api/categories/{id}
		switch r.Method {
		case http.MethodPut:
			handlers.UpdateCategory(w, r)
		case http.MethodDelete:
			handlers.DeleteCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return serveMux
}
