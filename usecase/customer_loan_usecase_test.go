package usecase

import (
	"loan-engine/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockLoans = []domain.Loan{
	{ID: 1, PrincipalAmount: 1000000, InterestRate: 10, WeekDuration: 4},
}

var mockCustomers = []domain.Customer{
	{ID: 1, Name: "John Doe"},
}

var mockCustomerLoans = []domain.CustomerLoan{
	{
		ID: 1,
	},
}

func Test_CalculateTotalRepayment(t *testing.T) {
	t.Run("Valid calculate total repayment", func(t *testing.T) {
		res := calculateTotalRepayment(domain.Loan{
			PrincipalAmount: 1000000,
			InterestRate:    10,
			WeekDuration:    4,
		})

		assert.Equal(t, 1100000, res)
	})
}

func Test_CreateCustomerLoan(t *testing.T) {
	originalLoans := loans
	originalCustomers := customers
	originalCustomerLoans := customerLoans

	defer func() {
		loans = originalLoans
		customers = originalCustomers
		customerLoans = originalCustomerLoans
	}()

	loans = mockLoans
	customers = mockCustomers
	customerLoans = []domain.CustomerLoan{}

	t.Run("Successfully create customer loan", func(t *testing.T) {
		input := domain.CustomerLoan{
			ID:         1,
			CustomerID: 1,
			LoanID:     1,
		}

		result, schedules := CreateCustomerLoan(input)

		assert.NotEmpty(t, result)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, 1, result.CustomerID)
		assert.Equal(t, 1, result.LoanID)
		assert.Equal(t, mockCustomers[0], result.Customer)
		assert.Equal(t, mockLoans[0], result.Loan)
		assert.Equal(t, 1100000, result.TotalRepayment)

		assert.Len(t, schedules, 4)
		for i, schedule := range schedules {
			assert.Equal(t, i, schedule.ID)
			assert.Equal(t, i+1, schedule.Week)
			assert.Equal(t, 275000, schedule.WeeklyBillingAmount)
			assert.Equal(t, domain.LoanBilingScheduleStatusUnpaid, schedule.Status)
			assert.True(t, schedule.DueDate.After(time.Now().AddDate(0, 0, 7*i-1)))
			assert.True(t, schedule.DueDate.Before(time.Now().AddDate(0, 0, 7*i+1)))
		}
	})

	t.Run("Return empty when loan not found", func(t *testing.T) {
		input := domain.CustomerLoan{
			ID:         2,
			CustomerID: 1,
			LoanID:     999,
		}

		result, schedules := CreateCustomerLoan(input)

		assert.Empty(t, result)
		assert.Empty(t, schedules)
	})

	t.Run("Return empty when customer not found", func(t *testing.T) {
		input := domain.CustomerLoan{
			ID:         2,
			CustomerID: 999,
			LoanID:     1,
		}

		result, schedules := CreateCustomerLoan(input)

		assert.Empty(t, result)
		assert.Empty(t, schedules)
	})
}

func Test_GetCustomerLoanByID(t *testing.T) {
	originalCustomerLoans := customerLoans

	defer func() {
		customerLoans = originalCustomerLoans
	}()

	customerLoans = mockCustomerLoans

	t.Run("Able to get customer loan by id", func(t *testing.T) {
		customerLoan, err := GetCustomerLoanByID(1)

		assert.NotEmpty(t, customerLoan)
		assert.Empty(t, err)
	})

	t.Run("Return empty customer loan and error message", func(t *testing.T) {
		customerLoans, err := GetCustomerLoanByID(2)

		assert.Empty(t, customerLoans)
		assert.NotEmpty(t, err)
	})
}
