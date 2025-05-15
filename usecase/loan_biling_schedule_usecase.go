package usecase

import (
	"loan-engine/domain"
	"time"
)

var loanBilingSchedules = []domain.LoanBilingSchedule{}

// TODO: possible not used
func GetAllLoanBilingSchedules() []domain.LoanBilingSchedule {
	return loanBilingSchedules
}

func CreateLoanBilingSchedule(customerLoan domain.CustomerLoan) []domain.LoanBilingSchedule {
	weeklyBillingAmount := customerLoan.TotalRepayment / customerLoan.Loan.WeekDuration
	for i := 0; i < customerLoan.Loan.WeekDuration; i++ {
		loanBilingSchedule := domain.LoanBilingSchedule{
			ID:                  i,
			CustomerLoan:        customerLoan,
			Week:                i + 1,
			WeeklyBillingAmount: weeklyBillingAmount,
			Status:              domain.LoanBilingScheduleStatusUnpaid,
			DueDate:             time.Now().AddDate(0, 0, 7*i),
		}
		loanBilingSchedules = append(loanBilingSchedules, loanBilingSchedule)
	}
	return loanBilingSchedules
}

func isCustomerMatch(customerLoan domain.CustomerLoan, loanBilingSchedule domain.LoanBilingSchedule) bool {
	return loanBilingSchedule.CustomerLoan.Customer.ID == customerLoan.Customer.ID
}

func isLoanMatch(customerLoan domain.CustomerLoan, loanBilingSchedule domain.LoanBilingSchedule) bool {
	return loanBilingSchedule.CustomerLoan.Loan.ID == customerLoan.Loan.ID
}

func isOnTimeOrLate(paymentDate time.Time, loanBilingSchedule domain.LoanBilingSchedule) bool {
	return loanBilingSchedule.DueDate.Before(paymentDate) || loanBilingSchedule.DueDate.Equal(paymentDate)
}

func PayLoanBilingSchedule(paymentDate time.Time, customerLoan domain.CustomerLoan) []domain.LoanBilingSchedule {
	for i, loanBilingSchedule := range loanBilingSchedules {
		if isOnTimeOrLate(paymentDate, loanBilingSchedule) &&
			isCustomerMatch(customerLoan, loanBilingSchedule) &&
			isLoanMatch(customerLoan, loanBilingSchedule) {
			loanBilingSchedules[i].Status = domain.LoanBilingScheduleStatusPaid
		}
	}
	return loanBilingSchedules
}

func GetLoanBillingOutStanding(customerLoan domain.CustomerLoan) int {
	totalOutStanding := 0
	for _, loanBilingSchedule := range loanBilingSchedules {
		if loanBilingSchedule.Status == domain.LoanBilingScheduleStatusUnpaid {
			totalOutStanding += loanBilingSchedule.WeeklyBillingAmount
		}
	}
	return totalOutStanding
}
