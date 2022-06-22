package init

import (
	"context"
	"database/sql"
	domain "eurovision/pkg/domain"
	"time"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var initCountries = []domain.Country{
	{
		Name: "Italy",
		Flag: "🇮🇹",
	},
	{
		Name: "France",
		Flag: "🇫🇷",
	},
	{
		Name: "Germany",
		Flag: "🇩🇪",
	},
	{
		Name: "Spain",
		Flag: "🇪🇸",
	},
	{
		Name: "United Kingdom",
		Flag: "🇬🇧",
	},
	{
		Name: "Albania",
		Flag: "🇦🇱",
	},
	{
		Name: "Latvia",
		Flag: "🇱🇻",
	},
	{
		Name: "Lithuania",
		Flag: "🇱🇹",
	},
	{
		Name: "Switzerland",
		Flag: "🇨🇭",
	},
	{
		Name: "Slovenia",
		Flag: "🇸🇮",
	},
	{
		Name: "Ukrain",
		Flag: "🇺🇦",
	},
	{
		Name: "Bulgaria",
		Flag: "🇧🇬",
	},
	{
		Name: "Netherlands",
		Flag: "🇳🇱",
	},
	{
		Name: "Moldova",
		Flag: "🇲🇩",
	},
	{
		Name: "Portugal",
		Flag: "🇵🇹",
	},
	{
		Name: "Croatia",
		Flag: "🇭🇷",
	},
	{
		Name: "Denmark",
		Flag: "🇩🇰",
	},
	{
		Name: "Austria",
		Flag: "🇦🇹",
	},
	{
		Name: "Iceland",
		Flag: "🇮🇸",
	},
	{
		Name: "Greece",
		Flag: "🇬🇷",
	},
	{
		Name: "Norway",
		Flag: "🇳🇴",
	},
	{
		Name: "Armenia",
		Flag: "🇦🇲",
	},
	{
		Name: "Finland",
		Flag: "🇫🇮",
	},
	{
		Name: "Israel",
		Flag: "🇮🇱",
	},
	{
		Name: "Serbia",
		Flag: "🇷🇸",
	},
	{
		Name: "Azerbaijan",
		Flag: "🇦🇿",
	},
	{
		Name: "Georgia",
		Flag: "🇬🇪",
	},
	{
		Name: "Malta",
		Flag: "🇲🇹",
	},
	{
		Name: "San Marino",
		Flag: "🇸🇲",
	},
	{
		Name: "Australia",
		Flag: "🇦🇺",
	},
	{
		Name: "Cyprus",
		Flag: "🇨🇾",
	},
	{
		Name: "Ireland",
		Flag: "🇮🇪",
	},
	{
		Name: "North Macedonia",
		Flag: "🇲🇰",
	},
	{
		Name: "Estonia",
		Flag: "🇪🇪",
	},
	{
		Name: "Romania",
		Flag: "🇷🇴",
	},
	{
		Name: "Poland",
		Flag: "🇵🇱",
	},
	{
		Name: "Montenegro",
		Flag: "🇲🇪",
	},
	{
		Name: "Belgium",
		Flag: "🇧🇪",
	},
	{
		Name: "Sweden",
		Flag: "🇸🇪",
	},
	{
		Name: "Czech Republic",
		Flag: "🇨🇿",
	},
}

const (
	username = "eurovision"
	password = "P,PO)+{l4!C{ff"
	hostname = "127.0.0.1:3306"
	dbname   = "eurovision"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname+" CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("Rows affected %d\n", no)

	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

func CreateCountriesTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS country;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("%d tables were dropped", res)

	query = `CREATE TABLE country(uuid VARCHAR(191) NOT NULL, name VARCHAR(191) NOT NULL, bandName VARCHAR(191), songName VARCHAR(191), flag BLOB, participating BOOLEAN NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating countries table: %d", rows)
	return nil
}

func AddCountries(db *sql.DB) error {
	query := "INSERT INTO country(uuid, name, bandName, songName, flag, participating) VALUES (?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	for _, country := range initCountries {
		newId, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Error %s when creating new UUID", err)
			return err
		}

		res, err := stmt.ExecContext(ctx, newId, country.Name, "", "", country.Flag, false)
		if err != nil {
			log.Printf("Error %s when inserting row into countries table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%s %s created %d time", country.Flag, country.Name, rows)
	}

	return nil
}
