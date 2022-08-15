package migrations

import (
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateCountriesTable(db *sqlx.DB) {
	query := `DROP TABLE IF EXISTS country;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating country table %s", err)
	}
	log.Println("Country table was dropped ⬇")

	query = `CREATE TABLE IF NOT EXISTS country(
				uuid char(36) NOT NULL, 
				name VARCHAR(191) NOT NULL, 
				slug VARCHAR(191) NOT NULL, 
				bandName VARCHAR(191), 
				songName VARCHAR(191), 
				flag BLOB NOT NULL, 
				participating BOOLEAN NOT NULL,
				UNIQUE (slug)) 
			  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating country table %s", err)
	}

	log.Println("Country table was created 😃")
}

func AddCountries(db *sqlx.DB) {
	query := "INSERT INTO country(uuid, name, slug, bandName, songName, flag, participating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	for _, country := range initCountries {
		id := uuid.New()

		_, err := db.Exec(query, id, country.Name, country.Slug, "", "", country.Flag, false)

		if err != nil {
			log.Fatalf("Error when inserting %s %s into country table: %s", country.Name, country.Flag, err)
		}

		log.Printf("%s %s created", country.Name, country.Flag)
	}
}
