package database

import (
	"expense-tracker/types"
	"log"
	"time"
)

func GetExpenses() ([]types.MonthlyBreakdown, error) {
	rows, err := db.Query("SELECT create_date, category, amount, note FROM expenses")
	if err != nil {
		log.Printf("Error querying expenses: %v", err)
		return nil, err
	}
	defer rows.Close()

	var expenses []types.MonthlyBreakdown
	for rows.Next() {
		var expense types.MonthlyBreakdown
		if err := rows.Scan(&expense.Date, &expense.Category, &expense.Amount, &expense.Note); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		parsedDate, err := time.Parse(time.RFC3339, expense.Date)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			return nil, err
		}
		expense.Date = parsedDate.Format("2006-01-02") // Format the date to YYYY-MM-DD

		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return expenses, nil
}
