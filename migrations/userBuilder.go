package migrations

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
)

func CreateUsersTable(db *sqlx.DB) bool {
	var count resultCount

	//query := `DROP TABLE IF EXISTS user;`
	//_, err := db.Exec(query)
	//if err != nil {
	//	log.Fatalf("Error %s when dropping user table", err)
	//}
	//log.Printf("User table was dropped ⬇")

	query := `CREATE TABLE IF NOT EXISTS user(
				uuid char(36) NOT NULL PRIMARY KEY, 
				name VARCHAR(191) NOT NULL,
				slug VARCHAR(191) NOT NULL UNIQUE, 
				authLvl TINYINT DEFAULT 0, 
				icon VARCHAR(191) DEFAULT 'default',
				KEY (slug));`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating user table %s", err)
	}

	log.Printf("User table was created 😃")

	query = `SELECT count(*) as count FROM user;`
	err = db.Select(&count, query)
	if err != nil {
		log.Fatal("Error when counting countries.", err)
	}

	return count[0] > 0
}

func AddUsers(db *sqlx.DB) {

	userQuery := "INSERT INTO user(uuid, name, slug, authLvl) VALUES (?, ?, ?, ?)"
	authQuery := "INSERT INTO auth(authToken, userId, authTokenExp, authLvl, slug) VALUES (?, ?, NOW() + INTERVAL 5 DAY, ?, ?)"

	for _, user := range initUsers {
		id := uuid.New()

		_, err := db.Exec(userQuery, id, user.Name, user.Slug, user.AuthLvl)
		if err != nil {
			log.Fatalf("User %s was not created. %s", user.Name, err)
		}

		switch user.AuthLvl {
		case enum.ADMIN:
			initAuth.GenerateSecureToken(40)
			_, err = db.Exec(authQuery, initAuth.AuthToken, id, user.AuthLvl, user.Slug)
			if err != nil {
				log.Fatalf("Authentication for user %s was not created. %s", user.Name, err)
			}
			//log.Printf("http://%s:%s/login/%s/%s", conf.App.Domain, conf.App.FrontendPort, initAuth.AuthToken, id)
			log.Printf("%s%s/login/%s/%s", conf.App.HttpProto, conf.App.Domain, initAuth.AuthToken, id)
			log.Printf("User %s created 👨‍💻", user.Name)
		case enum.BOT:
			conf.App.SetBotId(id)
			log.Printf("User %s created 🤖", user.Name)
		default:
			log.Printf("User %s created 👨", user.Name)
		}

	}
}
