package company

import (
	"log"

	"github.com/IsaErtunga/companies/internal/ports"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ca Adapter) Create(name string) (string, error) {
	log.Println("CREATING COMPANY")
	return "123", nil
}

func (ca Adapter) GetByID(id string) (ports.Company, error) {
	log.Println("GET BY ID")
	company := ports.Company{
		ID:   "123",
		Name: "Isas vård",
	}
	return company, nil
}

func (ca Adapter) Update(id string, name string) error {
	log.Println("UPDATE")
	return nil
}

func (ca Adapter) Delete(id string, name string) error {
	log.Println("Delete")
	return nil
}

// 	ListAll() ([]Company, error)
