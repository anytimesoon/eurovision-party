package main

import (
	"eurovision/conf"
	"eurovision/migrations"
	"eurovision/pkg/router"
	"log"
)

func main() {
	log.Println("Loading configuration 📃")
	conf.LoadConfig()
	log.Println("Config loaded ✅")

	log.Println("Starting application")
	db := migrations.Start()
	log.Println("Database migrations complete ✅")

	log.Println("Starting server 🖥")
	router.StartServer(&db)

	err := db.Close()
	if err != nil {
		log.Fatal("Failed to close db connection")
	}
	log.Println("Application closed")
}
