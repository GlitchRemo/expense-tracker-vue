package database

import (
	"expense-tracker/types"
	"log"
	"time"
)

func GetExpenses(id int) ([]types.MonthlyBreakdown, error) {
	rows, err := db.Query("SELECT id, create_date, category, amount, note, user_id FROM expenses WHERE user_id = $1 ORDER BY create_date DESC", id)
	if err != nil {
		log.Printf("Error querying expenses: %v", err)
		return nil, err
	}
	defer rows.Close()

	var expenses []types.MonthlyBreakdown
	for rows.Next() {
		var expense types.MonthlyBreakdown
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.Category, &expense.Amount, &expense.Note, &expense.UserID); err != nil {
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

func AddExpense(expense types.MonthlyBreakdown) error {
	query := `INSERT INTO expenses (create_date, category, amount, note, user_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, expense.Date, expense.Category, expense.Amount, expense.Note, expense.UserID)
	if err != nil {
		log.Printf("Error adding expense: %v", err)
	}
	return err
}

func DeleteExpense(id int) error {
	query := `DELETE FROM expenses WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting expense with id %d: %v", id, err)
	}
	return err
}

func UpdateExpense(expense types.MonthlyBreakdown) error {
	query := `UPDATE expenses SET create_date = $1, category = $2, amount = $3, note = $4 WHERE id = $5`
	_, err := db.Exec(query, expense.Date, expense.Category, expense.Amount, expense.Note, expense.ID)
	if err != nil {
		log.Printf("Error updating expense with id %d: %v", expense.ID, err)
	}
	return err
}
