package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

type CountryService interface {
	GetAllCountries() ([]dto.Country, *errs.AppError)
	UpdateCountry([]byte) (*dto.Country, *errs.AppError)
	SingleCountry(string) (*dto.Country, *errs.AppError)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func NewCountryService(repo domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repo}
}

func (service DefaultCountryService) GetAllCountries() ([]dto.Country, *errs.AppError) {
	countryData, err := service.repo.FindAllCountries()
	if err != nil {
		return nil, err
	}

	countriesDTO := make([]dto.Country, 0)
	for _, country := range countryData {
		countriesDTO = append(countriesDTO, country.ToDto())
	}

	return countriesDTO, nil
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
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError("Unable to read request")
	}

	country, appErr := service.repo.UpdateCountry(countryDTO)
	if err != nil {
		log.Println("FAILED to update country", err)
		return nil, appErr
	}

	countryDTO = country.ToDto()
	return &countryDTO, nil
}
