package db

import (
	"database/sql"
	initial "eurovision/pkg/init"
	"log"
)

var Conn *sql.DB

func Connect() {
	sqlDb, err := initial.Connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}

	log.Printf("Successfully connected to database")

	err = initial.CreateCountriesTable(sqlDb)
	if err != nil {
		log.Printf("Create country table failed with error %s", err)
		return
	}

	err = initial.AddCountries(sqlDb)
	if err != nil {
		log.Printf("Adding countries failed with error %s", err)
		return
	}

	Conn = sqlDb
}
