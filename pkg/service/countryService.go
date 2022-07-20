package service

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
)

type CountryService interface {
	GetAllCountries() (dto.Countries, error)
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
		countriesDTO.Data = append(countriesDTO.Data, dto.CountryData{ID: country.UUID, Flag: country.Flag, Name: country.Name, Slug: country.Slug, BandName: country.BandName, SongName: country.SongName, Participating: country.Participating})
	}
	countriesDTO.Success = true
	return countriesDTO, nil
}
