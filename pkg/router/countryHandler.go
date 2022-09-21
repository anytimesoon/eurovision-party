package router

import (
	"eurovision/pkg/service"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) FindAllCountries(resp http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.GetAllCountries()
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, countries)
	}
}

func (ch *CountryHandler) FindOneCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	country, err := ch.Service.SingleCountry(params["slug"])

	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, country)
	}
}

func (ch *CountryHandler) UpdateCountry(resp http.ResponseWriter, req *http.Request) {
	ok, appErr := currentSessions.authorize(req)
	if appErr != nil || !ok {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	country, appErr := ch.Service.UpdateCountry(body)

	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, country)
	}
}
