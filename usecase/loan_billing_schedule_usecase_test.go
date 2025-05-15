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
		DueDate: dueDate,
		Status:  domain.LoanBilingScheduleStatusUnpaid,
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
