package ports

import "github.com/IsaErtunga/companies/internal/core/domain"

type CompaniesRepository interface {
	GetById(id domain.CompanyID) (domain.Company, error)
	GetAll() ([]domain.Company, error)
	Insert(company domain.Company) error
}
