package handlers

import (
	"encoding/json"
	"net/http"

	"expense-tracker/types"

	"github.com/jmoiron/sqlx"
)

type AddExpenseRequest struct {
	Date     string         `json:"create_date"`
	Category types.Category `json:"category"`
	Amount   int            `json:"amount"`
	Note     string         `json:"note"`
}

func AddExpenseHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddExpenseRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO expenses (create_date, category, amount, note) VALUES ($1, $2, $3, $4)`
		_, err := db.Exec(query, req.Date, req.Category, req.Amount, req.Note)
		if err != nil {
			println("Error inserting expense:", err.Error())
			http.Error(w, "Failed to add expense", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Expense added successfully"))
	}
}
