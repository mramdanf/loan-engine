package usecase

import (
	"errors"
	"loan-engine/domain"
)

var customerLoans = []domain.CustomerLoan{}

func calculateTotalRepayment(loan domain.Loan) int {
	return loan.PrincipalAmount + (loan.PrincipalAmount * loan.InterestRate / 100)
}

func GetAllCustomerLoans() []domain.CustomerLoan {
	return customerLoans
}

func CreateCustomerLoan(customerLoan domain.CustomerLoan) (domain.CustomerLoan, []domain.LoanBillingSchedule) {
	loan, err := GetLoanByID(customerLoan.LoanID)
	if err != nil {
		return domain.CustomerLoan{}, []domain.LoanBillingSchedule{}
	}
	customer, err := GetCustomerByID(customerLoan.CustomerID)
	if err != nil {
		return domain.CustomerLoan{}, []domain.LoanBillingSchedule{}
	}
	customerLoan.Loan = loan
	customerLoan.Customer = customer
	customerLoan.TotalRepayment = calculateTotalRepayment(loan)

	customerLoans = append(customerLoans, customerLoan)

	loanBillingSchedules := CreateLoanBillingSchedule(customerLoan)

	return customerLoan, loanBillingSchedules
}

func GetCustomerLoanByID(id int) (domain.CustomerLoan, error) {
	for _, customerLoan := range customerLoans {
		if customerLoan.ID == id {
			return customerLoan, nil
		}
	}
	return domain.CustomerLoan{}, errors.New("customer loan not found")
}
