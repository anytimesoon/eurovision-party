package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateVotesTable(db *sqlx.DB) {
	query := `DROP TABLE IF EXISTS vote;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating vote table %s", err)
	}
	log.Printf("Vote table was dropped ⬇")

	query = `CREATE TABLE IF NOT EXISTS vote(
				uuid CHAR(36) NOT NULL, 
				userId CHAR(36) NOT NULL, 
				countryId CHAR(36) NOT NULL, 
				costume TINYINT, 
				song TINYINT, 
				performance TINYINT, 
				props TINYINT);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating vote table %s", err)
	}

	log.Printf("Vote table created 😃")
}
