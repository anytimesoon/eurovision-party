package migrations

import (
	"eurovision/conf"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Start(config *conf.App) sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn(config.DB))
	log.Println("Successfully connected to database")

	log.Println("Building tables 🏗")
	CreateAuthTable(sqlDb)
	CreateCountriesTable(sqlDb)
	CreateUsersTable(sqlDb)
	CreateCommentsTable(sqlDb)
	CreateVotesTable(sqlDb)

	log.Println("Seeding tables 🌱")
	AddCountries(sqlDb)
	AddUsers(sqlDb, config)

	return *sqlDb
}

func dsn(config conf.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Hostname, config.Port, config.DBName)
}
