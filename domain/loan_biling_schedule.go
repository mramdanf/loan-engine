package domain

import "time"

type LoanBilingScheduleStatus string

const (
	LoanBilingScheduleStatusUnpaid LoanBilingScheduleStatus = "unpaid"
	LoanBilingScheduleStatusPaid   LoanBilingScheduleStatus = "paid"
)

type LoanBilingSchedule struct {
	ID                  int                      `json:"id"`
	CustomerLoan        CustomerLoan             `json:"customer_loan"`
	Week                int                      `json:"week"`
	WeeklyBillingAmount int                      `json:"weekly_billing_amount"`
	Status              LoanBilingScheduleStatus `json:"status"`
	DueDate             time.Time                `json:"due_date"`
}
