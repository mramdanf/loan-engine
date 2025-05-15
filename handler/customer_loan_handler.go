package handler

import (
	"loan-engine/domain"
	"loan-engine/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCustomerLoan(c echo.Context) error {
	customerLoan := domain.CustomerLoan{}
	if err := c.Bind(&customerLoan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customerLoan, loanBilingSchedules := usecase.CreateCustomerLoan(customerLoan)
	return c.JSON(http.StatusOK, loanBilingSchedules)
}

func PayCustomerLoan(c echo.Context) error {
	var customerLoanPayment domain.CustomerLoanPayment
	if err := c.Bind(&customerLoanPayment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customerLoan, err := usecase.GetCustomerLoanByID(customerLoanPayment.CustomerLoanID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	loanBilingSchedules := usecase.PayLoanBilingSchedule(customerLoanPayment.PaymentDate, customerLoan)
	return c.JSON(http.StatusOK, loanBilingSchedules)
}
