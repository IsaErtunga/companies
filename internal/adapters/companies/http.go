package companiesadpt

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/IsaErtunga/companies/internal/core/ports"
	"github.com/go-chi/chi"
)

type Adapter struct {
	companiesService ports.CompaniesService
}

func NewHTTPAdapter(companiesService ports.CompaniesService) *Adapter {
	return &Adapter{
		companiesService: companiesService,
	}
}

func (adpt Adapter) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/", adpt.GetCompany)
	r.Post("/", adpt.CreateCompany)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", adpt.GetCompany)
	})

	return r
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("read error")
		return
	}

	var companyData struct {
		ID   string `json:"id`
		Name string `json:"name"`
		City string `json:"city"`
	}

	err = json.Unmarshal(body, &companyData)
	if err != nil {
		log.Println("[ERROR] deserializing", err)
	}

	err = adpt.companiesService.Create(companyData.Name)
	if err != nil {
		log.Println("create error")
		return
	}
}
