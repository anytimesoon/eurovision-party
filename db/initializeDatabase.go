package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var Conn *sql.DB

func Start() {
	sqlDb, err := connect()
	if err != nil {
		log.Panicf("Error %s when getting db connection", err)
	}

	log.Printf("Successfully connected to database")

	err = CreateCountriesTable(sqlDb)
	if err != nil {
		log.Panicf("FAILED to create country table %s", err)
	}

	err = AddCountries(sqlDb)
	if err != nil {
		log.Panicf("FAILED to add countries %s", err)
	}

	err = CreateUsersTable(sqlDb)
	if err != nil {
		log.Panicf("FAILED to create user table %s", err)
	}

	err = AddAdminUser(sqlDb)
	if err != nil {
		log.Panicf("FAILED to add admin user %s", err)
	}

	err = CreateCommentsTable(sqlDb)
	if err != nil {
		log.Panicf("FAILED to create comment table %s", err)
	}

	Conn = sqlDb
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", Username, Password, Hostname, DBName)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", DBName)
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

	db, err = sql.Open("mysql", dsn(DBName))
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
	log.Printf("Connected to DB %s successfully\n", DBName)
	return db, nil
}
