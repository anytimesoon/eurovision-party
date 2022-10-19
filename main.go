package main

import (
	"eurovision/conf"
	"eurovision/migrations"
	"eurovision/pkg/router"
	"log"
)

var appConf conf.App

func init() {
	appConf = conf.StartTui()
}

func main() {
	log.Println("Starting configuration 📃")
	appConf = conf.Setup()

	log.Println("Starting application")
	db := migrations.Start(appConf.DB)
	log.Println("Database migrations complete 🎉")

	log.Println("Starting server 🖥")
	router.StartServer(&db, appConf)

	db.Close()
	log.Println("Application closed")
}
