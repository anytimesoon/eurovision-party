package dao

import (
	db "eurovision/db"
	domain "eurovision/pkg/domain"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
)

func Countries() ([]domain.Country, error) {
	var countries []domain.Country

	stmt, err := db.Conn.Prepare("SELECT * FROM country")
	if err != nil {
		fmt.Println("stmt FAILED!")
		return countries, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("rows FAILED!")
		return countries, err
	}
	defer rows.Close()

	var id uuid.UUID
	var name string
	var bandName string
	var songName string
	var flag string
	var participating bool
	for rows.Next() {
		err = rows.Scan(&id, &name, &bandName, &songName, &flag, &participating)
		if err != nil {
			fmt.Println("scan FAILED!")
			return countries, err
		}
		countries = append(countries, domain.Country{UUID: id, Name: name, Flag: flag})
	}

	return countries, nil
}

func CountriesUpdate(country domain.Country) (domain.Country, error) {

	stmt, err := db.Conn.Prepare(
		"UPDATE country" +
			"SET participating = " + strconv.FormatBool(country.Participating) +
			" WHERE uuid = " + country.UUID.String() +
			" LIMIT 1;")
	if err != nil {
		log.Println("stmt FAILED!", stmt, err)
		country.Participating = !country.Participating
		return country, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("rows FAILED!")
		country.Participating = !country.Participating
		return country, err
	}
	defer rows.Close()

	return country, nil
}
