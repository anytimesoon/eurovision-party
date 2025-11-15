package api

import (
	"io"
	"log"
	"net/http"

	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/gorilla/mux"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) GetAllCountries(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	var countries *[]dto2.Country
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
	country := &dto2.Country{}
	if req.Context().Value("auth").(dto2.Auth).AuthLvl == authLvl.ADMIN {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println("FAILED to read body of COUNTRY UPDATE.", err)
			return
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
