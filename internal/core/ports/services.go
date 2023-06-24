package ports

import "github.com/IsaErtunga/companies/internal/core/domain"

type CompaniesService interface {
	Get(id domain.CompanyID) (domain.Company, error)
	Create(name string) (domain.CompanyID, error)
	// Update(id domain.CompanyID, name string) error
	// Delete(id domain.CompanyID) error
	// List() ([]domain.Company, error)
}
