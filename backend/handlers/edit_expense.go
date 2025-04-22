package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"expense-tracker/database"
	"expense-tracker/types"
)

func EditExpenseHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var updatedExpense types.MonthlyBreakdown
	if err := json.NewDecoder(r.Body).Decode(&updatedExpense); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedExpense.ID = id
	if err := database.UpdateExpense(updatedExpense); err != nil {
		http.Error(w, "Failed to update expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense updated successfully"))
}
