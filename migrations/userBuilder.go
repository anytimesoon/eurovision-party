package migrations

import (
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateUsersTable(db *sqlx.DB) {
	query := `DROP TABLE IF EXISTS user;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error %s when dropping user table", err)
	}
	log.Printf("User table was dropped ⬇")

	query = `CREATE TABLE IF NOT EXISTS user(
				uuid char(36) NOT NULL, 
				name VARCHAR(191) NOT NULL,
				email VARCHAR(191) NOT NULL, 
				slug VARCHAR(191) NOT NULL, 
				authLvl TINYINT DEFAULT 0, 
				icon VARCHAR(191) DEFAULT '/img/static/img/newuser.png',
				UNIQUE (slug, email));`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating user table %s", err)
	}

	log.Printf("User table was created 😃")
}

func AddUsers(db *sqlx.DB) {
	query := "INSERT INTO user(uuid, name, email, slug, authLvl) VALUES (?, ?, ?, ?, ?)"

	for _, user := range initUsers {
		id := uuid.New()

		_, err := db.Exec(query, id, user.Name, user.Email, user.Slug, user.AuthLvl)
		if err != nil {
			log.Fatalf("User %s was not created. %s", user.Name, err)
		}

		switch user.AuthLvl {
		case 1:
			log.Printf("User %s created 👨‍💻", user.Name)
		case 2:
			log.Printf("User %s created 🤖", user.Name)
		default:
			log.Printf("User %s created 👨", user.Name)
		}

	}
}
