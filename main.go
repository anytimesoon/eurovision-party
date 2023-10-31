package main

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/migrations"
	"github.com/anytimesoon/eurovision-party/pkg/router"
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
