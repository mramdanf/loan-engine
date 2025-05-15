package domain

type Loan struct {
	ID              int `json:"id"`
	PrincipalAmount int `json:"principal_amount"`
	InterestRate    int `json:"interest_rate"`
	WeekDuration    int `json:"week_duration"`
}
