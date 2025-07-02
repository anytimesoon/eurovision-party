package api

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) GetAllCountries(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	var countries *[]dto.Country
	countries, err = ch.Service.GetAllCountries()

	if err != nil {
		WriteResponse(resp, err.Code, *countries, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *countries, "")
	}
}

func (ch *CountryHandler) GetOneCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	country, err := ch.Service.GetOneCountry(params["slug"])

	if err != nil {
		WriteResponse(resp, err.Code, *country, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *country, "")
	}
}

func (ch *CountryHandler) GetParticipatingCountries(resp http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.GetParticipatingCountries()

	if err != nil {
		WriteResponse(resp, err.Code, *countries, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *countries, "")
	}
}

func (ch *CountryHandler) UpdateCountry(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	country := &dto.Country{}
	if req.Context().Value("auth").(dto.Auth).AuthLvl == authLvl.ADMIN {
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
		WriteResponse(resp, appErr.Code, *country, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *country, "")
	}
}
