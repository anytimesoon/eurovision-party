package dao

import (
	"context"
	db "eurovision/db"
	domain "eurovision/pkg/domain"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

var countryID uuid.UUID
var countryName string
var countrySlug string
var bandName string
var songName string
var flag string
var participating bool

func Countries() ([]domain.Country, error) {
	var countries []domain.Country
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	stmt, err := db.Conn.PrepareContext(ctx, "SELECT * FROM country")
	if err != nil {
		fmt.Println("FAILED to build query!")
		return countries, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println("rows FAILED!")
		return countries, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&countryID, &countryName, &bandName, &songName, &flag, &participating)
		if err != nil {
			log.Println("scan FAILED!")
			return countries, err
		}
		countries = append(countries, domain.Country{UUID: countryID, Name: countryName, Flag: flag})
	}

	return countries, nil
}

func Country(country domain.Country) (domain.Country, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`SELECT * FROM country WHERE slug = '%s'`, country.Slug)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return country, err
	}

	row := stmt.QueryRowContext(ctx)

	err = row.Scan(&countryID, &countryName, &countrySlug, &bandName, &songName, &flag, &participating)
	if err != nil {
		log.Printf("FAILED to scan because %s", err)
		return country, err
	}

	return domain.Country{UUID: countryID, Name: countryName, Slug: countrySlug, BandName: bandName, SongName: songName, Flag: flag, Participating: participating}, nil
}

func CountriesUpdate(country domain.Country, receivedCountry domain.Country) (domain.Country, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`UPDATE country SET bandName = '%s', songName = '%s', participating = %t WHERE uuid = '%s'`, receivedCountry.BandName, receivedCountry.SongName, receivedCountry.Participating, receivedCountry.UUID.String())
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return country, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! %s was not updated %s", country.Name, err)
		return country, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Printf("Error %s when finding rows affected", err)
		return country, err
	}

	log.Println("Country rows affected:", rowsAffected)
	return receivedCountry, nil
}
