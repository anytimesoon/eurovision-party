package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"log"
)

type CountryService interface {
	GetAllCountries() (dto.Countries, error)
	UpdateCountry([]byte) (dto.Country, error)
	SingleCountry(string) (dto.Country, error)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func NewCountryService(repo domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repo}
}

func (service DefaultCountryService) GetAllCountries() (dto.Countries, error) {
	var countriesDTO dto.Countries

	countryData, err := service.repo.FindAllCountries()
	if err != nil {
		countriesDTO.Message = err.Error()
		return countriesDTO, err
	}

	for _, country := range countryData {
		countriesDTO.Data = append(countriesDTO.Data, country.ToDto())
	}
	countriesDTO.Success = true
	return countriesDTO, nil
}

func (service DefaultCountryService) SingleCountry(slug string) (dto.Country, error) {
	var countryDTO dto.Country
	country, err := service.repo.FindOneCountry(slug)
	if err != nil {
		countryDTO.Message = err.Error()
		return countryDTO, err
	}

	return country.ToDto(), nil
}

func (service DefaultCountryService) UpdateCountry(body []byte) (dto.Country, error) {
	var countryDTO dto.Country
	err := json.Unmarshal(body, &countryDTO.Data)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		countryDTO.Message = err.Error()
		return countryDTO, err
	}

	country, err := service.repo.UpdateCountry(countryDTO.Data)
	if err != nil {
		log.Println("FAILED to update country", err)
		countryDTO.Message = err.Error()
		return countryDTO, err
	}

	return country.ToDto(), nil
}
