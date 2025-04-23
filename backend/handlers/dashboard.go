package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	// Extract userID from request (assuming it's passed as a query parameter for simplicity)
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	// Convert userID to integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	// Fetch total income and expense using the database function
	totalIncome, totalExpense, err := database.GetUserIncomeAndExpense(id)
	if err != nil {
		http.Error(w, "Failed to fetch user income and expense", http.StatusInternalServerError)
		return
	}

	// Fetch expenses (existing functionality)
	expenses, err := database.GetExpenses()
	if err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	// Prepare response data
	data := types.DashboardData{
		TotalIncome:      totalIncome,
		TotalExpenses:    totalExpense,
		MonthlyBreakdown: expenses,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
