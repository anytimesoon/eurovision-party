package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func StartMigrations() sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn())

	log.Printf("Successfully connected to database")

	CreateCountriesTable(sqlDb)

	AddCountries(sqlDb)

	CreateUsersTable(sqlDb)

	AddAdminUser(sqlDb)

	CreateCommentsTable(sqlDb)

	CreateVotesTable(sqlDb)

	return *sqlDb
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", Username, Password, Hostname, DBName)
}
