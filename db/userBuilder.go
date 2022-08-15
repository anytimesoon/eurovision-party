package db

import (
	"eurovision/pkg/domain"
	"fmt"
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
				slug VARCHAR(191) NOT NULL, 
				authLvl TINYINT DEFAULT 0, 
				icon VARCHAR(191) DEFAULT '/img/static/img/newuser.png',
				UNIQUE (slug));`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error when creating user table %s", err)
	}

	log.Printf("User table was created 😃")
}

func AddAdminUser(db *sqlx.DB) {
	adminUser := domain.User{
		UUID:    uuid.New(),
		Name:    "admin",
		Slug:    "admin",
		AuthLvl: domain.Admin,
	}

	query := fmt.Sprintf(`INSERT INTO user(uuid, name, slug, authLvl) VALUES ('%s', '%s', '%s', )`, adminUser.UUID.String(), adminUser.Name, adminUser.Slug)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("sql execution FAILED! admin user was not created. %s", err)
	}

	log.Println("Admin user created 👨‍💼")
}
