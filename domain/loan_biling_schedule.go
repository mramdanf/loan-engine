package domain

import "time"

type LoanBillingScheduleStatus string

const (
	LoanBillingScheduleStatusUnpaid LoanBillingScheduleStatus = "unpaid"
	LoanBillingScheduleStatusPaid   LoanBillingScheduleStatus = "paid"
)

type LoanBillingSchedule struct {
	ID                  int                       `json:"id"`
	CustomerLoan        CustomerLoan              `json:"customer_loan"`
	Week                int                       `json:"week"`
	WeeklyBillingAmount int                       `json:"weekly_billing_amount"`
	Status              LoanBillingScheduleStatus `json:"status"`
	DueDate             time.Time                 `json:"due_date"`
}
