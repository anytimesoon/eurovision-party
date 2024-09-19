package service

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"log"
)

type CountryService interface {
	GetAllCountries() (*[]dto.Country, *errs.AppError)
	UpdateCountry([]byte) (*dto.Country, *errs.AppError)
	SingleCountry(string) (*dto.Country, *errs.AppError)
	Participating() (*[]dto.Country, *errs.AppError)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func NewCountryService(repo domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repo}
}

func (service DefaultCountryService) GetAllCountries() (*[]dto.Country, *errs.AppError) {
	countryData, err := service.repo.FindAllCountries()
	if err != nil {
		return nil, err
	}

	return countriesToDto(*countryData), nil
}

func (service DefaultCountryService) SingleCountry(slug string) (*dto.Country, *errs.AppError) {
	country, err := service.repo.FindOneCountry(slug)
	if err != nil {
		return nil, err
	}

	countryDTO := country.ToDto()

	return &countryDTO, nil
}

func (service DefaultCountryService) UpdateCountry(body []byte) (*dto.Country, *errs.AppError) {
	var countryDTO dto.Country
	err := json.Unmarshal(body, &countryDTO)
	if err != nil {
		log.Printf("Failed to unmarshal json %s", err)
		return nil, errs.NewUnexpectedError("Unable to read request")
	}

	country, appErr := service.repo.UpdateCountry(countryDTO)
	if appErr != nil {
		return nil, appErr
	}

	countryDTO = country.ToDto()
	return &countryDTO, nil
}

func (service DefaultCountryService) Participating() (*[]dto.Country, *errs.AppError) {
	countryData, err := service.repo.FindParticipating()
	if err != nil {
		return nil, err
	}

	return countriesToDto(*countryData), nil
}

func countriesToDto(countryData []domain.Country) *[]dto.Country {
	countriesDTO := make([]dto.Country, 0)
	for _, country := range countryData {
		countriesDTO = append(countriesDTO, country.ToDto())
	}

	return &countriesDTO
}
