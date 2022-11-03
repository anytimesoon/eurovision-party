package main

import (
	"eurovision/conf"
	"eurovision/migrations"
	"eurovision/pkg/router"
	"log"
	"os"
)

var appConf conf.App

func init() {
	appConf = conf.StartTui()
}

func main() {
	log.Println("Starting application")
	db, err := migrations.Start(appConf.DB)
	if err != nil {
		log.Println("Your DB configuration was incorrect", err)
		os.Exit(1)
	}

	log.Println("Database migrations complete 🎉")

	log.Println("Starting server 🖥")
	router.StartServer(db, appConf)

	db.Close()
	log.Println("Application closed")
}
