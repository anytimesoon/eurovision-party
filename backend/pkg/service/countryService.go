package service

import (
	"encoding/json"
	"log"

	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
)

type CountryService interface {
	GetAllCountries() (*[]dto.Country, *errs.AppError)
	UpdateCountry([]byte) (*dto.Country, *errs.AppError)
	GetOneCountry(string) (*dto.Country, *errs.AppError)
	GetParticipatingCountries() (*[]dto.Country, *errs.AppError)
}

type DefaultCountryService struct {
	repo data.CountryRepository
}

func NewCountryService(repo data.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repo}
}

func (service DefaultCountryService) GetAllCountries() (*[]dto.Country, *errs.AppError) {
	countryData, err := service.repo.GetAllCountries()
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return countriesToDto(*countryData), nil
}

func (service DefaultCountryService) GetOneCountry(slug string) (*dto.Country, *errs.AppError) {
	country, err := service.repo.GetOneCountry(slug)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
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

	country, err := service.repo.UpdateCountry(dao.Country{}.FromDTO(countryDTO))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "country")
	}

	countryDTO = country.ToDto()
	return &countryDTO, nil
}

func (service DefaultCountryService) GetParticipatingCountries() (*[]dto.Country, *errs.AppError) {
	countryData, err := service.repo.GetParticipatingCountries()
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return countriesToDto(*countryData), nil
}

func countriesToDto(countryData []dao.Country) *[]dto.Country {
	countriesDTO := make([]dto.Country, 0)
	for _, country := range countryData {
		countriesDTO = append(countriesDTO, country.ToDto())
	}

	return &countriesDTO
}
