package main

import (
	"eurovision/conf"
	"eurovision/migrations"
	"eurovision/pkg/routes"
	"log"
)

func main() {
	log.Println("Starting configuration 📃")
	appConf := conf.Setup()

	log.Println("Starting application")
	db := migrations.Start(appConf.DB)
	log.Println("Database migrations complete 🎉")

	log.Println("Starting server 🖥")
	routes.StartServer(&db, appConf)

	db.Close()
	log.Println("Application closed")
}
