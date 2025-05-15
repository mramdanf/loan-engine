package usecase

import (
	"errors"
	"loan-engine/domain"
)

var loans = []domain.Loan{
	{ID: 1, PrincipalAmount: 5000000, InterestRate: 10, WeekDuration: 3},
}

func GetAllLoans() []domain.Loan {
	return loans
}

func GetLoanByID(id int) (domain.Loan, error) {
	for _, loan := range loans {
		if loan.ID == id {
			return loan, nil
		}
	}
	return domain.Loan{}, errors.New("loan not found")
}
