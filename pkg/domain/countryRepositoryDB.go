package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/jmoiron/sqlx"
	"github.com/timshannon/bolthold"
	"log"
)

type CountryRepositoryDb struct {
	client *sqlx.DB
	store  *bolthold.Store
}

func NewCountryRepositoryDb(db *sqlx.DB, store *bolthold.Store) CountryRepositoryDb {
	return CountryRepositoryDb{db, store}
}

func (db CountryRepositoryDb) FindAllCountries() (*[]Country, *errs.AppError) {
	countries := make([]Country, 0)

	var q bolthold.Query

	err := db.store.Find(&countries, &q)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &countries, nil
}

func (db CountryRepositoryDb) FindOneCountry(slug string) (*Country, *errs.AppError) {
	var country Country

	err := db.store.Get(slug, &country)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return &country, nil
}

func (db CountryRepositoryDb) UpdateCountry(countryDTO dto.Country) (*Country, *errs.AppError) {
	var country Country

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

func (db CountryRepositoryDb) FindParticipating() (*[]Country, *errs.AppError) {
	countries := make([]Country, 0)

	err := db.store.Find(&countries, bolthold.Where("Participating").Eq(true))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return &countries, nil
}
