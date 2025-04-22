package handlers

import (
	"net/http"
	"strconv"

	"expense-tracker/database"
)

func DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := database.DeleteExpense(id); err != nil {
		http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense deleted successfully"))
}
