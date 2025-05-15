package domain

import "time"

type CustomerLoan struct {
	ID             int      `json:"id"`
	CustomerID     int      `json:"customer_id" validate:"required"`
	LoanID         int      `json:"loan_id" validate:"required"`
	Loan           Loan     `json:"loan"`
	Customer       Customer `json:"customer"`
	TotalRepayment int      `json:"total_repayment"`
}

type CustomerLoanPayment struct {
	CustomerLoanID int       `json:"customer_loan_id" validate:"required"`
	PaymentDate    time.Time `json:"payment_date" validate:"required"`
}
