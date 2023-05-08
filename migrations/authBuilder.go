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
				authToken VARCHAR(191) NOT NULL, 
				authTokenExp DATETIME NOT NULL,
				sessionToken VARCHAR(191) DEFAULT '',
				sessionTokenExp DATETIME DEFAULT NOW(),
				authLvl TINYINT,
				slug VARCHAR(191));`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating auth table %s", err)
	}

	log.Printf("Auth table created 😃")
}
