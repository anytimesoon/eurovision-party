package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"log"
)

type CountryService interface {
	GetAllCountries() ([]dto.Country, error)
	UpdateCountry([]byte) (dto.Country, error)
	SingleCountry(string) (dto.Country, error)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func NewCountryService(repo domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repo}
}

func (service DefaultCountryService) GetAllCountries() ([]dto.Country, error) {
	countriesDTO := make([]dto.Country, 0)

	countryData, err := service.repo.FindAllCountries()
	if err != nil {
		return countriesDTO, err
	}

	for _, country := range countryData {
		countriesDTO = append(countriesDTO, country.ToDto())
	}

	return countriesDTO, nil
}

func (service DefaultCountryService) SingleCountry(slug string) (dto.Country, error) {
	var countryDTO dto.Country
	country, err := service.repo.FindOneCountry(slug)
	if err != nil {
		return countryDTO, err
	}

	return country.ToDto(), nil
}

func (service DefaultCountryService) UpdateCountry(body []byte) (dto.Country, error) {
	var countryDTO dto.Country
	err := json.Unmarshal(body, &countryDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return countryDTO, err
	}

	country, err := service.repo.UpdateCountry(countryDTO)
	if err != nil {
		log.Println("FAILED to update country", err)
		return countryDTO, err
	}

	return country.ToDto(), nil
}
