package dao

import (
	db "eurovision/db"
	domain "eurovision/pkg/domain"
	"fmt"
	"log"

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
	log.Printf("Country in: %+v \n", country)

	query := fmt.Sprintf(`UPDATE country SET participating = %t WHERE uuid = '%s'`, country.Participating, country.UUID.String())
	res, err := db.Conn.Exec(query)
	if err != nil {
		log.Println("sql execution FAILED!", country.Name, "was not updated", err)
		country.Participating = !country.Participating
		return country, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Println("FAILED to affect rows!", err)
		country.Participating = !country.Participating
		return country, err
	}

	log.Println("Country rows affected:", rowsAffected)
	return country, nil
}
