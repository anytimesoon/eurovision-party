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
	userQuery := "INSERT INTO user(uuid, name, email, slug, authLvl) VALUES (?, ?, ?, ?, ?)"
	authQuery := "INSERT INTO auth(token, userId, expiration, authLvl, slug) VALUES (?, ?, NOW() + INTERVAL 5 DAY, ?, ?)"

	for _, user := range initUsers {
		id := uuid.New()

		_, err := db.Exec(userQuery, id, user.Name, user.Email, user.Slug, user.AuthLvl)
		if err != nil {
			log.Fatalf("User %s was not created. %s", user.Name, err)
		}

		switch user.AuthLvl {
		case 1:
			initAuth.GenerateSecureToken()
			_, err = db.Exec(authQuery, initAuth.Token, id, user.AuthLvl, user.Slug)
			if err != nil {
				log.Fatalf("Authentication for user %s was not created. %s", user.Name, err)
			}
			log.Printf("http://localhost:8080/login/%s/%s", initAuth.Token, id)
			log.Printf("User %s created 👨‍💻", user.Name)
		case 2:
			log.Printf("User %s created 🤖", user.Name)
		default:
			log.Printf("User %s created 👨", user.Name)
		}

	}
}
