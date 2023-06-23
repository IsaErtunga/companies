package ports

type CompaniesPort interface {
	Create(name string) (string, error)
	GetByID(id string) (Company, error)
	Update(id string, name string) error
	Delete(id string) error
	// ListAll() ([]Company, error)
}

// extend with type CompanyID string

type Company struct {
	ID   string
	Name string
}
