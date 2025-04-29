package handlers

import (
	"fmt"
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
	fmt.Println("DashboardHandler called", r.URL.Query())
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
	fmt.Println("calling to database.GetUserIncomeAndExpense")
	totalIncome, totalExpense, err := database.GetUserIncomeAndExpense(id)
	if err != nil {
		fmt.Println("Error fetching user income and expense:", err)
		http.Error(w, "Failed to fetch user income and expense", http.StatusInternalServerError)
		return
	}

	// Fetch expenses (existing functionality)
	expenses, err := database.GetExpenses(id)
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
