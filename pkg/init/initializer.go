package init

import (
	"context"
	"database/sql"
	"time"

	// domain "eurovision/pkg/domain"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var countryNames = [40]string{"Italy", "France", "Germany", "Spain", "United Kingdom", "Albania", "Latvia", "Lithuania", "Switzerland", "Slovenia", "Ukraine", "Bulgaria", "Netherlands", "Moldova", "Portugal", "Croatia", "Denmark", "Austria", "Iceland", "Greece", "Norway", "Armenia", "Finland", "Israel", "Serbia", "Azerbaijan", "Georgia", "Malta", "San Marino", "Australia", "Cyprus", "Ireland", "North Macedonia", "Estonia", "Romania", "Poland", "Montenegro", "Belgium", "Sweden", "Czech Republic"}

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
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
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
	query := `DROP TABLE IF EXISTS country;` // `CREATE TABLE country(uuid VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, bandName VARCHAR(255) NOT NULL, songName VARCHAR(255) NOT NULL, flag VARCHAR(255) NOT NULL, participating BOOLEAN NOT NULL);`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("%d tables were dropped", res)

	query = `CREATE TABLE country(uuid VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, bandName VARCHAR(255), songName VARCHAR(255), flag VARCHAR(255), participating BOOLEAN NOT NULL);`
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

	for _, countryName := range countryNames {
		newId, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Error %s when creating new UUID", err)
			return err
		}

		res, err := stmt.ExecContext(ctx, newId, countryName, "", "", "", false)
		if err != nil {
			log.Printf("Error %s when inserting row into countries table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%s created %d time", countryName, rows)
	}

	return nil
}
