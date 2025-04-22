package handlers

import (
	"encoding/json"
	"net/http"

	"expense-tracker/database"
	"expense-tracker/types"
)

type AddExpenseRequest struct {
	Date     string         `json:"date"`
	Category types.Category `json:"category"`
	Amount   int            `json:"amount"`
	Note     string         `json:"note"`
}

func AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var req AddExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	err := database.AddExpense(types.MonthlyBreakdown{
		Date:     req.Date,
		Category: req.Category,
		Amount:   req.Amount,
		Note:     req.Note,
	})
	
	if err != nil {
		http.Error(w, "Failed to add expense", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Expense added successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
