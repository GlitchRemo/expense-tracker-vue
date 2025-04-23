package database

import (
	"fmt"
)

func GetUserIncomeAndExpense(userID int) (float64, float64, error) {
	var totalIncome, totalExpense float64

	// 1. Get total_income from users table
	userQuery := `SELECT COALESCE(total_income, 0) FROM users WHERE id = $1`
	err := db.QueryRow(userQuery, userID).Scan(&totalIncome)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching user income: %w", err)
	}

	// 2. Calculate total_expense from expenses table
	expenseQuery := `SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id = $1`
	err = db.QueryRow(expenseQuery, userID).Scan(&totalExpense)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching user expense: %w", err)
	}

	return totalIncome, totalExpense, nil
}
