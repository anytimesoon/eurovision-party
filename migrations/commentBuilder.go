package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateCommentsTable(db *sqlx.DB) {
	//query := `DROP TABLE IF EXISTS comment;`
	//_, err := db.Exec(query)
	//if err != nil {
	//	log.Fatalf("Error when creating comment table %s", err)
	//}
	//log.Printf("Comment table was dropped ⬇")

	query := `CREATE TABLE IF NOT EXISTS comment(
				uuid CHAR(36) NOT NULL, 
				userId CHAR(36) NOT NULL, 
				text VARCHAR(191) NOT NULL,
				fileName VARCHAR(255) DEFAULT '', 
				replyTo CHAR(36),
				createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                PRIMARY KEY (uuid));`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating comment table %s", err)
	}

	log.Printf("Comment table created 😃")
}
