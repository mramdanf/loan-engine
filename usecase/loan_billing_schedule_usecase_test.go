package usecase

import (
	"loan-engine/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var dueDate = time.Date(2025, 5, 15, 21, 42, 6, 906818000, time.FixedZone("+07:00", 7*60*60))

var mockLoanBillingSchedules = []domain.LoanBilingSchedule{
	{
		CustomerLoan: domain.CustomerLoan{
			Customer: domain.Customer{
				ID: 1,
			},
			Loan: domain.Loan{
				ID: 2,
			},
		},
		DueDate:             dueDate,
		Status:              domain.LoanBilingScheduleStatusUnpaid,
		WeeklyBillingAmount: 100000,
	},
}

func Test_PayLoanBilingSchedule(t *testing.T) {
	originalSchedules := loanBilingSchedules

	defer func() {
		loanBilingSchedules = originalSchedules
	}()

	loanBilingSchedules = mockLoanBillingSchedules

	t.Run("Successfully change status to paid", func(t *testing.T) {
		loanBilingSchedules[0].Status = domain.LoanBilingScheduleStatusUnpaid
		scheds := PayLoanBilingSchedule(dueDate, domain.CustomerLoan{
			Customer: domain.Customer{
				ID: 1,
			},
			Loan: domain.Loan{
				ID: 2,
			},
		})

		assert.Equal(t, domain.LoanBilingScheduleStatusPaid, scheds[0].Status)
	})

	t.Run("No billing data found", func(t *testing.T) {
		loanBilingSchedules[0].Status = domain.LoanBilingScheduleStatusUnpaid
		scheds := PayLoanBilingSchedule(dueDate, domain.CustomerLoan{
			Customer: domain.Customer{
				ID: 10,
			},
			Loan: domain.Loan{
				ID: 10,
			},
		})

		assert.Equal(t, domain.LoanBilingScheduleStatusUnpaid, scheds[0].Status)
	})
}

func Test_GetLoanBillingOutStanding(t *testing.T) {
	originalSchedules := loanBilingSchedules

	defer func() {
		loanBilingSchedules = originalSchedules
	}()

	loanBilingSchedules = mockLoanBillingSchedules

	t.Run("Successfully get loan billing outstanding", func(t *testing.T) {
		outstanding := GetLoanBillingOutStanding(domain.CustomerLoan{
			Customer: domain.Customer{
				ID: 1,
			},
			Loan: domain.Loan{
				ID: 2,
			},
		})

		assert.Equal(t, 100000, outstanding)
	})
}

func Test_IsLoanDelinquent(t *testing.T) {
	originalSchedules := loanBilingSchedules

	defer func() {
		loanBilingSchedules = originalSchedules
	}()

	var customerLoan = domain.CustomerLoan{
		Customer: domain.Customer{
			ID: 1,
		},
		Loan: domain.Loan{
			ID: 2,
		},
	}

	t.Run("Successfully check loan delinquent", func(t *testing.T) {
		loanBilingSchedules = []domain.LoanBilingSchedule{}

		for i := 0; i < 2; i++ {
			loanBilingSchedules = append(loanBilingSchedules, domain.LoanBilingSchedule{
				CustomerLoan:        customerLoan,
				DueDate:             dueDate,
				Status:              domain.LoanBilingScheduleStatusUnpaid,
				WeeklyBillingAmount: 100000,
			})
		}
		delinquent := IsLoanDelinquent(customerLoan, dueDate.AddDate(0, 0, 1))

		assert.True(t, delinquent)
	})

	t.Run("Successfully check loan not delinquent", func(t *testing.T) {
		loanBilingSchedules = []domain.LoanBilingSchedule{}

		for i := 0; i < 2; i++ {
			loanBilingSchedules = append(loanBilingSchedules, domain.LoanBilingSchedule{
				CustomerLoan:        customerLoan,
				DueDate:             dueDate,
				Status:              domain.LoanBilingScheduleStatusPaid,
				WeeklyBillingAmount: 100000,
			})
		}

		delinquent := IsLoanDelinquent(customerLoan, dueDate)

		assert.False(t, delinquent)
	})
}
