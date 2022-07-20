package main

import (
	"database/sql"
	initializer "eurovision/db"
	"eurovision/pkg/routes"
	"log"
)

var db sql.DB

func main() {
	log.Println("Starting application")
	db := initializer.StartMigrations()
	log.Println("Database migrations complete")

	log.Println("Starting server")
	routes.StartServer(&db)

	db.Close()
	log.Println("Application closed")
}
