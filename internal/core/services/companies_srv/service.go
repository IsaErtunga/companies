package companiessrv

import (
	"errors"

	"github.com/IsaErtunga/companies/internal/core/domain"
	"github.com/IsaErtunga/companies/internal/core/ports"
	"github.com/google/uuid"
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
	company, err := srv.repo.GetById(id)
	if err != nil {
		return domain.Company{}, errors.New("get company from repository has failed")
	}

	return company, nil
}

func (srv service) Create(name string, city string) error {
	company := domain.Company{
		ID:   domain.CompanyID(uuid.New().String()),
		Name: name,
		City: city,
	}

	err := srv.repo.Insert(company)
	if err != nil {
		return errors.New("create company has failed")
	}

	return nil
}

func (srv service) List() ([]domain.Company, error) {
	companies, err := srv.repo.GetAll()
	if err != nil {
		return []domain.Company{}, errors.New("create company has failed")
	}

	return companies, nil
}
