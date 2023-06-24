package ports

import "github.com/IsaErtunga/companies/internal/core/domain"

type CompaniesRepository interface {
	Get(id domain.CompanyID) (domain.Company, error)
	Insert(company domain.Company) error
}
