package companieshdl

import (
	"log"
	"net/http"

	"github.com/IsaErtunga/companies/internal/core/ports"
)

type HTTPHandler struct {
	companiesService ports.CompaniesService
}

func NewHTTPHandler(companiesService ports.CompaniesService) *HTTPHandler {
	return &HTTPHandler{
		companiesService: companiesService,
	}
}

func (rh HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	company, err := rh.companiesService.Get("123")
	if err != nil {
		log.Println("error")
		return
	}

	log.Println(company.Name)
}
