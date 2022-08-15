package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func StartMigrations() sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn())
	log.Println("Successfully connected to database")

	log.Println("Building tables 🏗")
	CreateCountriesTable(sqlDb)
	CreateUsersTable(sqlDb)
	CreateCommentsTable(sqlDb)
	CreateVotesTable(sqlDb)

	log.Println("Seeding tables 🌱")
	AddCountries(sqlDb)
	AddUsers(sqlDb)

	return *sqlDb
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", Username, Password, Hostname, DBName)
}
