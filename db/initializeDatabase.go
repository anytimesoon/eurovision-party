package db

import (
	"context"
	"database/sql"
	data "eurovision/initialData"
	"fmt"
	"log"
	"strings"
	"time"
)

var Conn *sql.DB

func Start() {
	sqlDb, err := connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}

	log.Printf("Successfully connected to database")

	err = CreateCountriesTable(sqlDb)
	if err != nil {
		log.Printf("FAILED to create country table %s", err)
		return
	}

	err = AddCountries(sqlDb)
	if err != nil {
		log.Printf("FAILED to add countries %s", err)
		return
	}

	err = CreateUsersTable(sqlDb)
	if err != nil {
		log.Printf("FAILED to create user table %s", err)
		return
	}

	err = AddAdminUser(sqlDb)
	if err != nil {
		log.Printf("FAILED to add countries %s", err)
		return
	}

	Conn = sqlDb
}

func slugify(name string) string {
	splitName := strings.Split(name, " ")
	slug := strings.Join(splitName, "-")
	return slug
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", data.Username, data.Password, data.Hostname, data.DBName)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", data.DBName)
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("Rows affected %d\n", rows)

	db, err = sql.Open("mysql", dsn(data.DBName))
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
	log.Printf("Connected to DB %s successfully\n", data.DBName)
	return db, nil
}
