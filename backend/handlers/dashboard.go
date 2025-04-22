package handlers

import (
	"encoding/json"
	"net/http"

	"expense-tracker/database"
	"expense-tracker/types"
)

var Categories = types.Categories{
	Groceries: "Groceries",
	Transport: "Transport",
	Utilities: "Utilities",
	Dining:    "Dining",
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	expenses, err := database.GetExpenses()
	if err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	data := types.DashboardData{
		TotalIncome:      5000, // Placeholder, can be fetched from DB if needed
		TotalExpenses:    2000, // Placeholder, can be calculated from expenses
		MonthlyBreakdown: expenses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
