package usecase

import (
	"loan-engine/domain"
	"time"
)

var loanBillingSchedules = []domain.LoanBillingSchedule{}

func CreateLoanBillingSchedule(customerLoan domain.CustomerLoan) []domain.LoanBillingSchedule {
	weeklyBillingAmount := customerLoan.TotalRepayment / customerLoan.Loan.WeekDuration
	for i := 0; i < customerLoan.Loan.WeekDuration; i++ {
		loanBillingSchedule := domain.LoanBillingSchedule{
			ID:                  i,
			CustomerLoan:        customerLoan,
			Week:                i + 1,
			WeeklyBillingAmount: weeklyBillingAmount,
			Status:              domain.LoanBillingScheduleStatusUnpaid,
			DueDate:             time.Now().AddDate(0, 0, 7*i),
		}
		loanBillingSchedules = append(loanBillingSchedules, loanBillingSchedule)
	}
	return loanBillingSchedules
}

func isCustomerMatch(customerLoan domain.CustomerLoan, loanBillingSchedule domain.LoanBillingSchedule) bool {
	return loanBillingSchedule.CustomerLoan.Customer.ID == customerLoan.Customer.ID
}

func isLoanMatch(customerLoan domain.CustomerLoan, loanBillingSchedule domain.LoanBillingSchedule) bool {
	return loanBillingSchedule.CustomerLoan.Loan.ID == customerLoan.Loan.ID
}

func isPaymentDateAfterOrOnDueDate(paymentDate time.Time, loanBillingSchedule domain.LoanBillingSchedule) bool {
	return loanBillingSchedule.DueDate.Before(paymentDate) || loanBillingSchedule.DueDate.Equal(paymentDate)
}

func isBillingUnpaid(loanBillingSchedule domain.LoanBillingSchedule) bool {
	return loanBillingSchedule.Status == domain.LoanBillingScheduleStatusUnpaid
}

func isDueDateBeforeCurrentDate(loanBillingSchedule domain.LoanBillingSchedule, currentDate time.Time) bool {
	return loanBillingSchedule.DueDate.Before(currentDate)
}

func PayLoanBillingSchedule(paymentDate time.Time, customerLoan domain.CustomerLoan) []domain.LoanBillingSchedule {
	for i, loanBillingSchedule := range loanBillingSchedules {
		if isPaymentDateAfterOrOnDueDate(paymentDate, loanBillingSchedule) &&
			isCustomerMatch(customerLoan, loanBillingSchedule) &&
			isLoanMatch(customerLoan, loanBillingSchedule) {
			loanBillingSchedules[i].Status = domain.LoanBillingScheduleStatusPaid
		}
	}
	return loanBillingSchedules
}

func GetLoanBillingOutStanding(customerLoan domain.CustomerLoan) int {
	totalOutStanding := 0
	for _, loanBillingSchedule := range loanBillingSchedules {
		if isCustomerMatch(customerLoan, loanBillingSchedule) &&
			isLoanMatch(customerLoan, loanBillingSchedule) &&
			isBillingUnpaid(loanBillingSchedule) {
			totalOutStanding += loanBillingSchedule.WeeklyBillingAmount
		}
	}
	return totalOutStanding
}

func IsLoanDelinquent(customerLoan domain.CustomerLoan, currentDate time.Time) bool {
	unpaidCount := 0
	maxUnpaidCount := 2

	for _, loanBillingSchedule := range loanBillingSchedules {
		if isCustomerMatch(customerLoan, loanBillingSchedule) &&
			isLoanMatch(customerLoan, loanBillingSchedule) &&
			isBillingUnpaid(loanBillingSchedule) &&
			isDueDateBeforeCurrentDate(loanBillingSchedule, currentDate) &&
			unpaidCount < maxUnpaidCount {
			unpaidCount++
		}
	}
	return unpaidCount >= maxUnpaidCount
}
