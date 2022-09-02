package handler

import (
	"bytes"
	"encoding/json"
	"eurovision/mocks/service"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var countryRouter *mux.Router
var ch CountryHandler
var mockCountryService *service.MockCountryService
var mockCountries []dto.Country
var mockCountry dto.Country
var countryJSON []byte
var countryBody *bytes.Buffer

func setupCountryTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCountryService = service.NewMockCountryService(ctrl)
	ch = CountryHandler{mockCountryService}
	mockCountries = []dto.Country{
		{ID: uuid.New(), Name: "tEsTcOuNtRy", Slug: "testcountry", BandName: "WTF", SongName: "That's right", Flag: "🇫🇷", Participating: true},
		{ID: uuid.New(), Name: "bigDickia", Slug: "bigdickia", BandName: "Suck it", SongName: "Aw yeah", Flag: "🇩🇪", Participating: false},
	}

	mockCountry = mockCountries[0]
	countryJSON, _ = json.Marshal(mockCountry)
	countryBody = bytes.NewBuffer(countryJSON)

	countryRouter = mux.NewRouter()
	countryRouter.HandleFunc("/country", ch.FindAllCountries).Methods(http.MethodGet)
	countryRouter.HandleFunc("/country", ch.UpdateCountry).Methods(http.MethodPut)
	countryRouter.HandleFunc("/country/{slug}", ch.FindOneCountry).Methods(http.MethodGet)
}

func Test_all_countries_route_should_return_countries_with_200_code(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().GetAllCountries().Return(mockCountries, nil)

	req, _ := http.NewRequest(http.MethodGet, "/country", nil)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	countries := make([]dto.User, 0)
	json.Unmarshal(recorder.Body.Bytes(), &countries)

	if len(countries) != 2 {
		t.Error("Expected 2 countries, but found", len(countries))
	}
}

func Test_all_countries_route_should_return_500_code(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().GetAllCountries().Return(nil, errs.NewUnexpectedError("Couldn't find countries"))

	req, _ := http.NewRequest(http.MethodGet, "/country", nil)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_country_update_route_returns_500_code(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().UpdateCountry(countryJSON).Return(nil, errs.NewUnexpectedError("Couldn't update country"))

	req, _ := http.NewRequest(http.MethodPut, "/country", countryBody)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_country_update_route_returns_updated_country(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().UpdateCountry(countryJSON).Return(&mockCountry, nil)

	req, _ := http.NewRequest(http.MethodPut, "/country", countryBody)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	var country dto.Country
	json.Unmarshal(recorder.Body.Bytes(), &country)

	if country != mockCountry {
		t.Errorf("Expected %+v to equal %+v", country, mockCountry)
	}
}

func Test_find_one_country_route_returns_500_code(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().SingleCountry("testcountry").Return(nil, errs.NewUnexpectedError("Couldn't find country"))

	req, _ := http.NewRequest(http.MethodGet, "/country/testcountry", nil)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_find_one_country_route_returns_user(t *testing.T) {
	setupCountryTest(t)

	mockCountryService.EXPECT().SingleCountry("testcountry").Return(&mockCountry, nil)

	req, _ := http.NewRequest(http.MethodGet, "/country/testcountry", nil)

	recorder := httptest.NewRecorder()
	countryRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	var returnedCountry dto.Country
	_ = json.Unmarshal(recorder.Body.Bytes(), &returnedCountry)

	if returnedCountry != mockCountry {
		t.Errorf("Expected %+v to equal %+v", returnedCountry, mockCountry)
	}
}
