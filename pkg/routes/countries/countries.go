package countries

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/domain"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func All(writer http.ResponseWriter, req *http.Request) {
	countries, err := dao.Countries()
	if err != nil {
		log.Println("home FAILED!")
		return
	}

	json.NewEncoder(writer).Encode(countries)
}

func FindOne(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	countryName := params["name"]

	partialCountry := domain.Country{Name: countryName}
	country, err := dao.Country(partialCountry)
	if err != nil {
		log.Printf("FAILED to find %s", countryName)
	}

	json.NewEncoder(writer).Encode(country)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var receivedCountry domain.Country

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COUNTRY UPDATE!")
		return
	}

	err = json.Unmarshal(body, &receivedCountry)
	if err != nil {
		log.Println("FAILED to unmarshal json!")
		return
	}

	country, err := dao.Country(receivedCountry)
	if err != nil {
		log.Printf("FAILED to find %s", receivedCountry.Name)
	}

	updatedCountry, err := dao.CountriesUpdate(country, receivedCountry)
	if err != nil {
		log.Printf("FAILED to update %s", receivedCountry.Name)
	}

	json.NewEncoder(writer).Encode(updatedCountry)
}
