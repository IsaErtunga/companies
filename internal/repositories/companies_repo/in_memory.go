package companiesrepo

import (
	"encoding/json"
	"errors"

	"github.com/IsaErtunga/companies/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo memkvs) GetById(id domain.CompanyID) (domain.Company, error) {
	// return domain.Company{ID: "123", Name: "isa", City: "ss"}, nil
	if value, ok := repo.kvs[string(id)]; ok {
		company := domain.Company{}
		err := json.Unmarshal(value, &company)
		if err != nil {
			return domain.Company{}, errors.New("fail to get value from kvs")
		}

		return company, nil
	}

	return domain.Company{}, errors.New("game not found in kvs")
}

func (repo memkvs) Insert(company domain.Company) error {
	companyJSON, err := json.Marshal(company)
	if err != nil {
		return errors.New("failed to serialize company")
	}

	repo.kvs[string(company.ID)] = companyJSON
	return nil
}

func (repo memkvs) GetAll() ([]domain.Company, error) {
	values := make([]domain.Company, 0, len(repo.kvs))
	for _, value := range repo.kvs {
		company := domain.Company{}
		err := json.Unmarshal(value, &company)
		if err != nil {
			return []domain.Company{}, errors.New("fail to parse JSON from kvs")
		}
		values = append(values, company)
	}
	return values, nil
}
