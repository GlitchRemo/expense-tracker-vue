package types

type Category string

type Categories struct {
	Groceries Category
	Transport Category
	Utilities Category
	Dining    Category
}

type MonthlyBreakdown struct {
	Date     string   `json:"date" db:"create_date"`
	Category Category `json:"category"`
	Amount   int      `json:"amount"`
	Note     string   `json:"note"`
}

type DashboardData struct {
	TotalIncome      int                `json:"totalIncome"`
	TotalExpenses    int                `json:"totalExpenses"`
	MonthlyBreakdown []MonthlyBreakdown `json:"monthlyBreakdown"`
}
