package entity

//Loan ..
type Loan struct {
	LoanID      int     `json:"loan_id" binding:"required,number"`
	CustomerID  int     `json:"custumer_id" binding:"required,number"`
	StartDate   string  `json:"start_date" `
	EndDate     string  `json:"end_date" binding:"required"`
	AmountLoan  float64 `json:"amount_loan" binding:"required,number"`
	Percent     float64 `json:"percent" binding:"required,number"`
	Lifetime    int     `json:"lifetime" binding:"required,number"`
	CurrentDebt float64 `json:"current_debt" binding:"required,number"`
}
