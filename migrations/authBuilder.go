package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateAuthTable(db *sqlx.DB) {
	query := `DROP TABLE IF EXISTS auth;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating auth table %s", err)
	}
	log.Printf("Auth table was dropped ⬇")

	query = `CREATE TABLE IF NOT EXISTS auth(
				userId CHAR(36) NOT NULL, 
				token VARCHAR(191) NOT NULL, 
				expiration DATETIME NOT NULL);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating auth table %s", err)
	}

	log.Printf("Auth table created 😃")
}
