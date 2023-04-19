package service

//
//import (
//	"encoding/json"
//	mockDomain "eurovision/mocks/domain"
//	"eurovision/pkg/domain"
//	"eurovision/pkg/dto"
//	"eurovision/pkg/errs"
//	"fmt"
//	"net/http"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	"github.com/google/uuid"
//)
//
//var countryService CountryService
//var mockCountryRepository *mockDomain.MockCountryRepository
//var mockCountries []domain.Country
//var mockCountry domain.Country
//var mockCountriesDTO []dto.Country
//var mockCountryDTO dto.Country
//var countryJSON []byte
//var invalidCountryDTO dto.Country
//var invalidCountryJSON []byte
//
//func setupCountryTest(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockCountryRepository = mockDomain.NewMockCountryRepository(ctrl)
//	countryService = DefaultCountryService{mockCountryRepository}
//	mockCountries = []domain.Country{
//		{UUID: uuid.New(), Name: "tEsTcOuNtRy", Slug: "testcountry", BandName: "WTF", SongName: "That's right", Flag: "🇫🇷", Participating: true},
//		{UUID: uuid.New(), Name: "bigDickia", Slug: "bigdickia", BandName: "Suck it", SongName: "Aw yeah", Flag: "🇩🇪", Participating: false},
//	}
//	mockCountry = mockCountries[0]
//
//	mockCountriesDTO = []dto.Country{
//		mockCountry.ToDto(),
//		mockCountries[1].ToDto(),
//	}
//	mockCountryDTO = mockCountriesDTO[0]
//	countryJSON, _ = json.Marshal(mockCountryDTO)
//
//	invalidCountryDTO = dto.Country{ID: uuid.New(), Name: "iVaLiDcOuNtRy", Slug: "invalidcountry", BandName: "", SongName: "", Flag: "🇫🇷", Participating: true}
//	invalidCountryJSON, _ = json.Marshal(invalidCountryDTO)
//}
//
//func Test_country_service_returns_all_countrys(t *testing.T) {
//	setupCountryTest(t)
//
//	mockCountryRepository.EXPECT().FindAllCountries().Return(mockCountries, nil)
//
//	// result, _ := countryService.GetAllCountries()
//
//	// if result[0] != mockCountriesDTO[0] || result[1] != mockCountriesDTO[1] {
//	// 	t.Error("Returned countrys do not match expected")
//	// }
//}
//
//func Test_all_countries_service_returns_500_error(t *testing.T) {
//	setupCountryTest(t)
//
//	mockCountryRepository.EXPECT().FindAllCountries().Return(nil, errs.NewUnexpectedError("DB error occurred"))
//
//	_, err := countryService.GetAllCountries()
//
//	if err.Code != http.StatusInternalServerError {
//		t.Errorf("Expected 500 error, but got %d", err.Code)
//	}
//}
//
//func Test_update_country_service_returns_updated_country(t *testing.T) {
//	setupCountryTest(t)
//
//	mockCountryRepository.EXPECT().UpdateCountry(mockCountryDTO).Return(&mockCountry, nil)
//
//	result, _ := countryService.UpdateCountry(countryJSON)
//
//	if result.ID != mockCountryDTO.ID {
//		t.Error("Returned countrys do not match expected")
//	}
//}
//
//func Test_update_country_service_returns_500_error(t *testing.T) {
//	setupCountryTest(t)
//
//	mockCountryRepository.EXPECT().UpdateCountry(mockCountryDTO).Return(nil, errs.NewUnexpectedError("DB error occurred"))
//
//	result, err := countryService.UpdateCountry(countryJSON)
//
//	if result != nil {
//		fmt.Printf("%+v", result)
//	}
//	if err.Code != http.StatusInternalServerError {
//		t.Errorf("Expected 500 error, but got %d", err.Code)
//	}
//}
//
//func Test_update_country_service_returns_400_error(t *testing.T) {
//	setupCountryTest(t)
//	_, err := countryService.UpdateCountry(invalidCountryJSON)
//
//	if err.Code != http.StatusBadRequest {
//		t.Errorf("Expected 400 error, but got %d", err.Code)
//	}
//}
//
//func Test_single_country_service_returns_one_country(t *testing.T) {
//	setupCountryTest(t)
//
//	mockSlug := "testcountry"
//
//	mockCountryRepository.EXPECT().FindOneCountry(mockSlug).Return(&mockCountry, nil)
//
//	result, _ := countryService.SingleCountry(mockSlug)
//
//	if result.ID != mockCountryDTO.ID {
//		t.Error("Returned countrys do not match expected")
//	}
//}
//
//func Test_single_country_service_returns_500_error(t *testing.T) {
//	setupCountryTest(t)
//
//	mockCountry := "testcountry"
//
//	mockCountryRepository.EXPECT().FindOneCountry(mockCountry).Return(nil, errs.NewUnexpectedError("DB error occurred"))
//
//	_, err := countryService.SingleCountry(mockCountry)
//
//	if err.Code != http.StatusInternalServerError {
//		t.Errorf("Expected 500 error, but got %d", err.Code)
//	}
//}
