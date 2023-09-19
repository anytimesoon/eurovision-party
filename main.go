package main

import (
	"eurovision/conf"
	"eurovision/migrations"
	"eurovision/pkg/router"
	"log"
)

func main() {
	log.Println("Starting configuration 📃")
	appConf := conf.Setup()

	log.Println("Starting application")
	db := migrations.Start(&appConf)
	log.Println("Database migrations complete 🎉")

	log.Println("Starting server 🖥")
	router.StartServer(&db, appConf)

	err := db.Close()
	if err != nil {
		log.Fatal("Failed to close db connection")
	}
	log.Println("Application closed")
}
