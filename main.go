package main

import (
	"eurovision/migrations"
	"eurovision/pkg/routes"
	"log"
)

func main() {
	log.Println("Starting application")
	db := migrations.Start()
	log.Println("Database migrations complete 🎉")

	log.Println("Starting server")
	routes.StartServer(&db)

	db.Close()
	log.Println("Application closed")
}
