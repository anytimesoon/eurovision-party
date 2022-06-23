package dao

import (
	db "eurovision/db"
	domain "eurovision/pkg/domain"
	"fmt"

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

	return countries, err
}
