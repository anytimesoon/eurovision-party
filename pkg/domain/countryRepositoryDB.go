package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CountryRepositoryDb struct {
	client *sqlx.DB
}

func NewCountryRepositoryDb(db *sqlx.DB) CountryRepositoryDb {
	return CountryRepositoryDb{db}
}

func (db CountryRepositoryDb) FindAllCountries() ([]Country, *errs.AppError) {
	countries := make([]Country, 0)

	query := "SELECT * FROM country"
	err := db.client.Select(&countries, query)
	if err != nil {
		log.Println("Error while querying country table", err)
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return countries, nil
}

func (db CountryRepositoryDb) FindOneCountry(slug string) (*Country, *errs.AppError) {
	var country Country

	query := fmt.Sprintf(`SELECT * FROM country WHERE slug = '%s'`, slug)

	err := db.client.Get(&country, query)
	if err != nil {
		log.Println("Error while selecting one country", err)
		return nil, errs.NewUnexpectedError("Country not found")
	}

	return &country, nil
}

func (db CountryRepositoryDb) UpdateCountry(countryDTO dto.Country) (*Country, *errs.AppError) {
	var country Country

	query := fmt.Sprintf(`UPDATE country SET bandName = '%s', songName = '%s', participating = %t WHERE uuid = '%s'`, countryDTO.BandName, countryDTO.SongName, countryDTO.Participating, countryDTO.ID.String())

	_, err := db.client.NamedExec(query, country)
	if err != nil {
		log.Println("Error while updating country table", err)
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	query = "SELECT * FROM country WHERE slug = ?"
	err = db.client.Get(&country, query, countryDTO.Slug)
	if err != nil {
		log.Println("Error while fetching country after update", err)
		return nil, errs.NewUnexpectedError("Country updated, but not found")
	}

	return &country, nil
}
