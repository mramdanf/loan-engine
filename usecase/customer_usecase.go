package usecase

import (
	"errors"
	"loan-engine/domain"
)

var customers = []domain.Customer{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Malcolm Doe"},
}

func GetAllCustomers() []domain.Customer {
	return customers
}

func GetCustomerByID(id int) (domain.Customer, error) {
	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return domain.Customer{}, errors.New("customer not found")
}
