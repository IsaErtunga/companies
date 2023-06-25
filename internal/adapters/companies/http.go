package companiesadpt

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/IsaErtunga/companies/internal/core/domain"
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
	r.Get("/", adpt.ListCompanies)

	r.Route("/{companyId}", func(r chi.Router) {
		r.Get("/", adpt.GetCompany)
	})

	return r
}

func (adpt Adapter) GetCompany(w http.ResponseWriter, r *http.Request) {
	companyId := domain.CompanyID(chi.URLParam(r, "companyId"))
	company, err := adpt.companiesService.Get(companyId)
	if err != nil {
		log.Println("get error")
		return
	}

	err = json.NewEncoder(w).Encode(company)
	if err != nil {
		log.Println("[ERROR] serializing", err)
	}
}

func (adpt Adapter) ListCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := adpt.companiesService.List()
	if err != nil {
		log.Println("list error")
		return
	}

	err = json.NewEncoder(w).Encode(companies)
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
		ID   string `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	err = json.Unmarshal(body, &companyData)
	if err != nil {
		log.Println("[ERROR] deserializing", err)
	}

	err = adpt.companiesService.Create(companyData.Name, companyData.City)
	if err != nil {
		log.Println("create error")
		return
	}
}
