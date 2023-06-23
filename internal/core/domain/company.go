package domain

type Company struct {
	ID   CompanyID `json:"id"`
	Name string    `json:"name"`
	City string    `json:"city"`
}

type CompanyID string
