package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/timshannon/bolthold"
	"log"
)

type CountryRepository interface {
	FindAllCountries() (*[]dao.Country, *errs.AppError)
	FindOneCountry(string) (*dao.Country, *errs.AppError)
	UpdateCountry(dto.Country) (*dao.Country, *errs.AppError)
	FindParticipating() (*[]dao.Country, *errs.AppError)
}

type CountryRepositoryDb struct {
	store *bolthold.Store
}

func NewCountryRepositoryDb(store *bolthold.Store) CountryRepositoryDb {
	return CountryRepositoryDb{store}
}

func (db CountryRepositoryDb) FindAllCountries() (*[]dao.Country, *errs.AppError) {
	countries := make([]dao.Country, 0)

	var q bolthold.Query

	err := db.store.Find(&countries, &q)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &countries, nil
}

func (db CountryRepositoryDb) FindOneCountry(slug string) (*dao.Country, *errs.AppError) {
	var country dao.Country

	err := db.store.Get(slug, &country)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return &country, nil
}

func (db CountryRepositoryDb) UpdateCountry(countryDTO dto.Country) (*dao.Country, *errs.AppError) {
	var country dao.Country

	country = country.FromDTO(countryDTO)
	//var countries []Country
	//_ = db.store.Find(&countries, &bolthold.Query{})

	err := db.store.Update(country.Slug, country)
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "country")
	}

	return &country, nil
}

func (db CountryRepositoryDb) FindParticipating() (*[]dao.Country, *errs.AppError) {
	countries := make([]dao.Country, 0)

	err := db.store.Find(&countries, bolthold.Where("Participating").Eq(true))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return &countries, nil
}
