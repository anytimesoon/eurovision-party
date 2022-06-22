package init

import (
	"context"
	"database/sql"
	domain "eurovision/pkg/domain"
	"time"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var initCountries = []domain.Country{
	domain.Country{
		Name: "Italy",
		Flag: "🇮🇹",
	},
	domain.Country{
		Name: "France",
		Flag: "🇫🇷",
	},
	domain.Country{
		Name: "Germany",
		Flag: "🇩🇪",
	},
	domain.Country{
		Name: "Spain",
		Flag: "🇪🇸",
	},
	domain.Country{
		Name: "United Kingdom",
		Flag: "🇬🇧",
	},
	domain.Country{
		Name: "Albania",
		Flag: "🇦🇱",
	},
	domain.Country{
		Name: "Latvia",
		Flag: "🇱🇻",
	},
	domain.Country{
		Name: "Lithuania",
		Flag: "🇱🇹",
	},
	domain.Country{
		Name: "Switzerland",
		Flag: "🇨🇭",
	},
	domain.Country{
		Name: "Slovenia",
		Flag: "🇸🇮",
	},
	domain.Country{
		Name: "Ukrain",
		Flag: "🇺🇦",
	},
	domain.Country{
		Name: "Bulgaria",
		Flag: "🇧🇬",
	},
	domain.Country{
		Name: "Netherlands",
		Flag: "🇳🇱",
	},
	domain.Country{
		Name: "Moldova",
		Flag: "🇲🇩",
	},
	domain.Country{
		Name: "Portugal",
		Flag: "🇵🇹",
	},
	domain.Country{
		Name: "Croatia",
		Flag: "🇭🇷",
	},
	domain.Country{
		Name: "Denmark",
		Flag: "🇩🇰",
	},
	domain.Country{
		Name: "Austria",
		Flag: "🇦🇹",
	},
	domain.Country{
		Name: "Iceland",
		Flag: "🇮🇸",
	},
	domain.Country{
		Name: "Greece",
		Flag: "🇬🇷",
	},
	domain.Country{
		Name: "Norway",
		Flag: "🇳🇴",
	},
	domain.Country{
		Name: "Armenia",
		Flag: "🇦🇲",
	},
	domain.Country{
		Name: "Finland",
		Flag: "🇫🇮",
	},
	domain.Country{
		Name: "Israel",
		Flag: "🇮🇱",
	},
	domain.Country{
		Name: "Serbia",
		Flag: "🇷🇸",
	},
	domain.Country{
		Name: "Azerbaijan",
		Flag: "🇦🇿",
	},
	domain.Country{
		Name: "Georgia",
		Flag: "🇬🇪",
	},
	domain.Country{
		Name: "Malta",
		Flag: "🇲🇹",
	},
	domain.Country{
		Name: "San Marino",
		Flag: "🇸🇲",
	},
	domain.Country{
		Name: "Australia",
		Flag: "🇦🇺",
	},
	domain.Country{
		Name: "Cyprus",
		Flag: "🇨🇾",
	},
	domain.Country{
		Name: "Ireland",
		Flag: "🇮🇪",
	},
	domain.Country{
		Name: "North Macedonia",
		Flag: "🇲🇰",
	},
	domain.Country{
		Name: "Estonia",
		Flag: "🇪🇪",
	},
	domain.Country{
		Name: "Romania",
		Flag: "🇷🇴",
	},
	domain.Country{
		Name: "Poland",
		Flag: "🇵🇱",
	},
	domain.Country{
		Name: "Montenegro",
		Flag: "🇲🇪",
	},
	domain.Country{
		Name: "Belgium",
		Flag: "🇧🇪",
	},
	domain.Country{
		Name: "Sweden",
		Flag: "🇸🇪",
	},
	domain.Country{
		Name: "Czech Republic",
		Flag: "🇨🇿",
	},
}

const (
	username = "eurovision"
	password = "P,PO)+{l4!C{ff"
	hostname = "127.0.0.1:3306"
	dbname   = "eurovision"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname+" CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("Rows affected %d\n", no)

	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

func CreateCountriesTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS country;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("%d tables were dropped", res)

	query = `CREATE TABLE country(uuid VARCHAR(191) NOT NULL, name VARCHAR(191) NOT NULL, bandName VARCHAR(191), songName VARCHAR(191), flag BLOB, participating BOOLEAN NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating countries table: %d", rows)
	return nil
}

func AddCountries(db *sql.DB) error {
	query := "INSERT INTO country(uuid, name, bandName, songName, flag, participating) VALUES (?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	for _, country := range initCountries {
		newId, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Error %s when creating new UUID", err)
			return err
		}

		res, err := stmt.ExecContext(ctx, newId, country.Name, "", "", country.Flag, false)
		if err != nil {
			log.Printf("Error %s when inserting row into countries table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%s %s created %d time", country.Flag, country.Name, rows)
	}

	return nil
}
