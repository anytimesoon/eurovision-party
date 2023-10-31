package migrations

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type resultCount []uint8

func Start() sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn())
	log.Println("Successfully connected to database")

	log.Println("Building tables 🏗")
	CreateAuthTable(sqlDb)
	hasCountries := CreateCountriesTable(sqlDb)
	hasUsers := CreateUsersTable(sqlDb)
	CreateCommentsTable(sqlDb)
	CreateVotesTable(sqlDb)

	log.Println("Seeding tables 🌱")
	if !hasCountries {
		AddCountries(sqlDb)
	}

	if !hasUsers {
		AddUsers(sqlDb)
	}

	return *sqlDb
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.App.DbUsername, conf.App.DbPassword, conf.App.DbHostname, conf.App.DbPort, conf.App.DbName)
}
