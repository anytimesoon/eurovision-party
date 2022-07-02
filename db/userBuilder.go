package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateUsersTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS user;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when dropping user table", err)
		return err
	}
	log.Println("User table was dropped")

	query = `CREATE TABLE user(uuid char(36) NOT NULL, name VARCHAR(191) NOT NULL, authLvl TINYINT DEFAULT 0, icon VARCHAR(191));`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating user table: %d", rows)
	return nil
}

func AddAdminUser(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := fmt.Sprintf(`INSERT INTO user(uuid, name, authLvl) VALUES ('%s', 'admin', 1)`, uuid.New().String())
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! admin user was not created. %s", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	log.Println("User rows affected:", rowsAffected)

	return nil
}
