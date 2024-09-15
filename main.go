package main

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/migrations"
	"github.com/anytimesoon/eurovision-party/pkg/router"
	"github.com/timshannon/bolthold"
	"log"
	"path/filepath"
)

func main() {
	log.Println("Starting Eurovision backend")
	log.Println("Loading configuration 📃")
	conf.LoadConfig()
	log.Println("Config loaded ✅")

	log.Println("Starting application")
	store, err := bolthold.Open(filepath.Join(conf.App.DbPath, "data.db"), 0600, nil)
	if err != nil {
		log.Fatal("Failed to open KV database")
	}

	defer func(store *bolthold.Store) {
		err := store.Close()
		if err != nil {
			log.Fatal("Failed to close KV database")
		}
	}(store)

	db := migrations.Start(store)
	log.Println("Database migrations complete ✅")

	log.Println("Starting server 🖥")
	router.StartServer(&db, store)

	err = db.Close()
	if err != nil {
		log.Fatal("Failed to close db connection")
	}
	log.Println("Application closed")
}
