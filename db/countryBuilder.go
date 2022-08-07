package db

import (
	"context"
	"eurovision/pkg/utils"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateCountriesTable(db *sqlx.DB) error {
	query := `DROP TABLE IF EXISTS country;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating country table", err)
		return err
	}
	log.Printf("%d table was dropped", res)

	query = `CREATE TABLE country(
				uuid char(36) NOT NULL, 
				name VARCHAR(191) NOT NULL, 
				slug VARCHAR(191) NOT NULL, 
				bandName VARCHAR(191), 
				songName VARCHAR(191), 
				flag BLOB NOT NULL, 
				participating BOOLEAN NOT NULL,
				UNIQUE (slug)) 
			  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating country table", err)
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

func AddCountries(db *sqlx.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "INSERT INTO country(uuid, name, slug, bandName, songName, flag, participating) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	for _, country := range InitCountries {
		slug := utils.Slugify(country.Name)
		newId, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Error %s when creating new UUID", err)
			return err
		}

		res, err := stmt.ExecContext(ctx, newId, country.Name, slug, "", "", country.Flag, false)
		if err != nil {
			log.Printf("Error %s when inserting row into countries table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%+v, created %d time", country, rows)
	}

	return nil
}
