package db

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateVotesTable(db *sqlx.DB) error {
	query := `DROP TABLE IF EXISTS vote;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating vote table", err)
		return err
	}
	log.Printf("%d table was dropped", res)

	query = `CREATE TABLE vote(uuid CHAR(36) NOT NULL, userId CHAR(36) NOT NULL, countryId CHAR(36) NOT NULL, costume TINYINT, song TINYINT, performance TINYINT, props TINYINT);`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating vote table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating vote table: %d", rows)
	return nil
}
