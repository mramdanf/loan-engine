package main

import (
	"loan-engine/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/customer-loan", handler.CreateCustomerLoan)
	e.POST("/customer-loan/payment", handler.PayCustomerLoan)
	e.GET("/customer-loan/:customer_loan_id/outstanding", handler.GetCustomerLoanOutStanding)
	e.GET("/customer-loan/:customer_loan_id/delinquent", handler.IsCustomerLoanDelinquent)

	e.Logger.Fatal(e.Start(":8080"))
}
