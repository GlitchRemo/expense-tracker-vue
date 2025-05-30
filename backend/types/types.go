package types

type MonthlyBreakdown struct {
	ID         int    `json:"id" db:"id"`
	Date       string `json:"date" db:"create_date"`
	CategoryID int    `json:"categoryId" db:"category_id"`
	Amount     int    `json:"amount"`
	Note       string `json:"note"`
	UserID     int    `json:"userId" db:"user_id"`
}

type DashboardData struct {
	TotalIncome      float64            `json:"totalIncome"`
	TotalExpenses    float64            `json:"totalExpenses"`
	MonthlyBreakdown []MonthlyBreakdown `json:"monthlyBreakdown"`
}
