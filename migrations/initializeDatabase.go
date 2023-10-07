package migrations

import (
	"eurovision/conf"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Start() sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn())
	log.Println("Successfully connected to database")

	log.Println("Building tables 🏗")
	CreateAuthTable(sqlDb)
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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.Db.Username, conf.Db.Password, conf.Db.Hostname, conf.Db.Port, conf.Db.Name)
}
