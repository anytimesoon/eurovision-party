package router

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"eurovision/pkg/errs"
	"eurovision/pkg/service"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) FindAllCountries(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	var countries *[]dto.Country
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin {
		countries, err = ch.Service.GetAllCountries()
	} else {
		err = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}
	if err != nil {
		writeResponse(resp, req, err.Code, countries, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, countries, "")
	}
}

func (ch *CountryHandler) FindOneCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	country, err := ch.Service.SingleCountry(params["slug"])

	if err != nil {
		writeResponse(resp, req, err.Code, country, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, country, "")
	}
}

func (ch *CountryHandler) Participating(resp http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.Participating()

	if err != nil {
		writeResponse(resp, req, err.Code, countries, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, countries, "")
	}
}

func (ch *CountryHandler) UpdateCountry(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var country *dto.Country
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		log.Println("Update request body:", body)
		country, appErr = ch.Service.UpdateCountry(body)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, country, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, country, "")
	}
}
