package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/timshannon/bolthold"
	"log"
)

type CountryRepository interface {
	GetAllCountries() (*[]dao.Country, error)
	GetOneCountry(string) (*dao.Country, error)
	UpdateCountry(dao.Country) (*dao.Country, error)
	GetParticipatingCountries() (*[]dao.Country, error)
}

type CountryRepositoryDb struct {
	store *bolthold.Store
}

func NewCountryRepositoryDb(store *bolthold.Store) CountryRepositoryDb {
	return CountryRepositoryDb{store}
}

func (db CountryRepositoryDb) GetAllCountries() (*[]dao.Country, error) {
	countries := make([]dao.Country, 0)

	err := db.store.Find(&countries, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying country table for all countries", err)
		return nil, err
	}

	return &countries, nil
}

func (db CountryRepositoryDb) GetOneCountry(slug string) (*dao.Country, error) {
	var country dao.Country

	err := db.store.Get(slug, &country)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (db CountryRepositoryDb) UpdateCountry(country dao.Country) (*dao.Country, error) {
	err := db.store.Update(country.Slug, country)
	if err != nil {
		log.Printf("Error while updating country %s. %s", country.Name, err)
		return nil, err
	}

	return &country, nil
}

func (db CountryRepositoryDb) GetParticipatingCountries() (*[]dao.Country, error) {
	countries := make([]dao.Country, 0)

	err := db.store.Find(&countries, bolthold.Where("Participating").Eq(true))
	if err != nil {
		return nil, err
	}

	return &countries, nil
}
