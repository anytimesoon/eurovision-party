package api

import (
	dto2 "github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
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
	var countries *[]dto2.Country
	countries, err = ch.Service.GetAllCountries()

	if err != nil {
		WriteResponse(resp, req, err.Code, *countries, err.Message)
	} else {
		WriteResponse(resp, req, http.StatusOK, *countries, "")
	}
}

func (ch *CountryHandler) FindOneCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	country, err := ch.Service.SingleCountry(params["slug"])

	if err != nil {
		WriteResponse(resp, req, err.Code, *country, err.Message)
	} else {
		WriteResponse(resp, req, http.StatusOK, *country, "")
	}
}

func (ch *CountryHandler) Participating(resp http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.Participating()

	if err != nil {
		WriteResponse(resp, req, err.Code, *countries, err.Message)
	} else {
		WriteResponse(resp, req, http.StatusOK, *countries, "")
	}
}

func (ch *CountryHandler) UpdateCountry(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var country *dto2.Country
	if req.Context().Value("auth").(dto2.Auth).AuthLvl == enum.ADMIN {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		log.Println("Update request body:", string(body))
		country, appErr = ch.Service.UpdateCountry(body)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, req, appErr.Code, *country, appErr.Message)
	} else {
		WriteResponse(resp, req, http.StatusOK, *country, "")
	}
}
