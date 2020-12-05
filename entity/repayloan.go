package entity

import "time"

//RePayLoan ..
type RePayLoan struct {
	RePayLoanID int       `json:"loanpay_id"`
	CustomerID  int       `json:"costumer_id"`
	LoanID      int       `json:"loan_id"`
	PayAmount   float64   `json:"pay_amount"`
	Date        time.Time `json:"date"`
}
