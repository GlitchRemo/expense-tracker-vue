package handlers

import (
	"encoding/json"
	"net/http"

	"expense-tracker/types"
)

var Categories = types.Categories{
	Groceries: "Groceries",
	Transport: "Transport",
	Utilities: "Utilities",
	Dining:    "Dining",
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := types.DashboardData{
		TotalIncome:   5000,
		TotalExpenses: 2000,
		MonthlyBreakdown: []types.MonthlyBreakdown{
			{Date: "2023-01-01", Category: Categories.Groceries, Amount: 400, Note: "This is a note for the expense"},
			{Date: "2023-01-01", Category: Categories.Transport, Amount: 400, Note: "This is a note for the expense"},
			{Date: "2023-02-01", Category: Categories.Utilities, Amount: 300, Note: "This is a note for the expense"},
			{Date: "2023-03-01", Category: Categories.Dining, Amount: 500, Note: "This is a note for the expense"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
