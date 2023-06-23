package ports

type AppPort interface {
	CreateCompany(name string, description string) (string, error)
	GetCompanyByID(id string) (Company, error)
	UpdateCompany(id string, name string, description string) error
	DeleteCompany(id string) error
	ListCompanies() ([]Company, error)
}
