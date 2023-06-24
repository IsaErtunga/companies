package companiesadpt

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IsaErtunga/companies/internal/core/ports"
)

type Adapter struct {
	companiesService ports.CompaniesService
}

func NewHTTPAdapter(companiesService ports.CompaniesService) *Adapter {
	return &Adapter{
		companiesService: companiesService,
	}
}

func (adpt Adapter) GetCompany(w http.ResponseWriter, r *http.Request) {
	company, err := adpt.companiesService.Get("123")
	if err != nil {
		log.Println("get error")
		return
	}

	err = json.NewEncoder(w).Encode(company)
	if err != nil {
		log.Println("[ERROR] serializing", err)
	}
}

func (adpt Adapter) CreateCompany(w http.ResponseWriter, r *http.Request) {
	_, err := adpt.companiesService.Create("temp")
	if err != nil {
		log.Println("create error")
		return
	}
}
