package handler

import (
	"encoding/json"
	"eurovision/pkg/service"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) FindAllCountries(resp http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.GetAllCountries()
	if err != nil {
		log.Println("Failed to get all countries", err)
	}

	json.NewEncoder(resp).Encode(countries)
}

func (ch *CountryHandler) FindOneCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	country, err := ch.Service.SingleCountry(params["slug"])
	if err != nil {
		log.Println("Failed to get single country", err)
	}

	json.NewEncoder(resp).Encode(country)
}

func (ch *CountryHandler) UpdateCountry(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COUNTRY UPDATE!", err)
		return
	}

	country, err := ch.Service.UpdateCountry(body)
	if err != nil {
		log.Println("Failed to update country", err)
	}

	json.NewEncoder(resp).Encode(country)
}
