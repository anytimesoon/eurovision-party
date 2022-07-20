package domain

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type CountryRepositoryDb struct {
	client *sqlx.DB
}

func NewCountryRepositoryDb(db *sqlx.DB) CountryRepositoryDb {
	return CountryRepositoryDb{db}
}

func (db CountryRepositoryDb) FindAllCountries() ([]Country, error) {
	countries := make([]Country, 0)

	query := "SELECT * FROM country"
	err := db.client.Select(&countries, query)
	if err != nil {
		log.Println("Error while querying country table", err)
		return nil, err
	}

	return countries, nil
}
