package db

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateCommentsTable(db *sqlx.DB) error {
	query := `DROP TABLE IF EXISTS comment;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating comment table", err)
		return err
	}
	log.Printf("%d table was dropped", res)

	query = `CREATE TABLE comment(uuid CHAR(36) NOT NULL, userId CHAR(36) NOT NULL, text VARCHAR(191) NOT NULL, createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating comment table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating comment table: %d", rows)
	return nil
}
