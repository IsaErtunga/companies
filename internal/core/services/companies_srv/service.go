package companiessrv

import (
	"errors"

	"github.com/IsaErtunga/companies/internal/core/domain"
	"github.com/IsaErtunga/companies/internal/core/ports"
)

type service struct {
	repo ports.CompaniesRepository
}

func New(repo ports.CompaniesRepository) *service {
	return &service{
		repo: repo,
	}
}

func (srv service) Get(id domain.CompanyID) (domain.Company, error) {
	company, err := srv.repo.Get(id)
	if err != nil {
		return domain.Company{}, errors.New("get company from repository has failed")
	}

	return company, nil
}

func (srv service) Create(name string) (domain.CompanyID, error) {
	company := domain.Company{
		ID:   "123",
		Name: "FIRST",
		City: "Sthlm",
	}
	err := srv.repo.Insert(company)
	if err != nil {
		return "", errors.New("create company has failed")
	}

	return company.ID, nil
}
