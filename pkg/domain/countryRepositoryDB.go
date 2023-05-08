package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"

	"github.com/jmoiron/sqlx"
)

type CountryRepositoryDb struct {
	client *sqlx.DB
}

func NewCountryRepositoryDb(db *sqlx.DB) CountryRepositoryDb {
	return CountryRepositoryDb{db}
}

func (db CountryRepositoryDb) FindAllCountries() (*[]Country, *errs.AppError) {
	countries := make([]Country, 0)

	query := "SELECT * FROM country"
	err := db.client.Select(&countries, query)
	if err != nil {
		log.Println("Error while querying country table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &countries, nil
}

func (db CountryRepositoryDb) FindOneCountry(slug string) (*Country, *errs.AppError) {
	var country Country

	query := "SELECT * FROM country WHERE slug = ?"

	err := db.client.Get(&country, query, slug)
	if err != nil {
		log.Println("Error while selecting one country", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	return &country, nil
}

func (db CountryRepositoryDb) UpdateCountry(countryDTO dto.Country) (*Country, *errs.AppError) {
	var country Country

	updateCountryQuery := "UPDATE country SET bandName = ?, songName = ?, participating = ? WHERE slug = ?"
	getCountryQuery := "SELECT * FROM country WHERE slug = ?"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error while beginning tx for country %s. %s", countryDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "country")
	}

	_, err = tx.Exec(updateCountryQuery, countryDTO.BandName, countryDTO.SongName, countryDTO.Participating, countryDTO.Slug)
	if err != nil {
		log.Printf("Error while updating country table for %s. %s", countryDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "country")
	}

	err = tx.Get(&country, getCountryQuery, countryDTO.Slug)
	if err != nil {
		log.Printf("Error while fetching country %s after update. %s", countryDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "country")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while committing tx for country %s. %s", countryDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "country")
	}

	return &country, nil
}

func (db CountryRepositoryDb) FindParticipating() (*[]Country, *errs.AppError) {
	countries := make([]Country, 0)

	query := "SELECT * FROM country WHERE participating = true"
	err := db.client.Select(&countries, query)
	if err != nil {
		log.Println("Error while querying country table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &countries, nil
}
