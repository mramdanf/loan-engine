package handler

import (
	"loan-engine/domain"
	"loan-engine/usecase"
	"net/http"
	"strconv"
	"time"

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
	parsedDate, err := time.Parse("2006-01-02 15:04:05", customerLoanPayment.PaymentDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format. Expected format: YYYY-MM-DD HH:mm:ss")
	}
	loanBilingSchedules := usecase.PayLoanBilingSchedule(parsedDate, customerLoan)
	return c.JSON(http.StatusOK, loanBilingSchedules)
}

func GetCustomerLoanOutStanding(c echo.Context) error {
	customerLoanID := c.Param("customer_loan_id")
	customerLoanIDInt, err := strconv.Atoi(customerLoanID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customerLoan, err := usecase.GetCustomerLoanByID(customerLoanIDInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	outStanding := usecase.GetLoanBillingOutStanding(customerLoan)
	return c.JSON(http.StatusOK, outStanding)
}

func IsCustomerLoanDelinquent(c echo.Context) error {
	customerLoanID := c.Param("customer_loan_id")
	currentDate := c.QueryParam("current_date")
	customerLoanIDInt, err := strconv.Atoi(customerLoanID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	parsedDate, err := time.Parse("2006-01-02 15:04:05", currentDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format. Expected format: YYYY-MM-DD HH:mm:ss")
	}

	customerLoan, err := usecase.GetCustomerLoanByID(customerLoanIDInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	isDelinquent := usecase.IsLoanDelinquent(customerLoan, parsedDate)
	return c.JSON(http.StatusOK, isDelinquent)
}
