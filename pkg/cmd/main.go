package main

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/timshannon/bolthold"
	"log"
	"path/filepath"
	"time"
)

func main() {
	log.Println("Starting Eurovision backend")
	log.Println("Loading configuration 📃")
	conf.LoadConfig()
	log.Println("Config loaded ✅")

	log.Println("Starting application")
	store, err := bolthold.Open(filepath.Join(conf.App.DbPath, "data.db"), 0600, nil)
	if err != nil {
		log.Fatal("Failed to open KV database: ", err)
	}

	defer func(store *bolthold.Store) {
		err := store.Close()
		if err != nil {
			log.Fatal("Failed to close KV database")
		}
	}(store)

	addCountries(store)
	addUsers(store)

	log.Println("Database migrations complete ✅")

	log.Println("Starting server 🖥")
	StartServer(store)

	log.Println("Application closed")
}

func addCountries(store *bolthold.Store) {
	for _, country := range initCountriesWithParticipating {
		err := store.Insert(country.Slug, country)
		if err != nil {
			log.Printf("Skipping %s %s: already exists in country table", country.Flag, country.Name)
		}
	}
}

func addUsers(store *bolthold.Store) {
	admins := make([]dao.User, 0)
	err := store.Find(
		&admins,
		bolthold.
			Where("AuthLvl").
			Eq(enum.ADMIN).
			Index("AuthLvl"),
	)
	if err != nil {
		log.Println("Error when finding admins:", err)
	}
	if len(admins) == 0 {
		err = store.Insert(initAdminUser.UUID.String(), initAdminUser)
		if err != nil {
			log.Printf("%s alread exists in user table", initAdminUser.Name)
		}

		adminAuth := dao.Auth{
			UserId:       initAdminUser.UUID,
			AuthTokenExp: time.Now().Add(time.Hour * 24 * 100),
			AuthLvl:      enum.ADMIN,
			Slug:         initAdminUser.Slug,
		}
		adminAuth.GenerateSecureToken(40)
		err = store.Insert(adminAuth.AuthToken, adminAuth)
		if err != nil {
			log.Fatal("Error when inserting admin auth token:", err)
		}

		for _, country := range initCountriesWithParticipating {
			err = store.Upsert(
				fmt.Sprintf("%s_%s", initAdminUser.UUID, country.Slug),
				dao.Vote{
					UserId:      initAdminUser.UUID,
					CountrySlug: country.Slug,
					Costume:     0,
					Song:        0,
					Performance: 0,
					Props:       0,
				},
			)
			if err != nil {
				log.Fatalf("Error while inserting vote into vote table during vote creation. %s", err)
			}
		}

		log.Printf("%s%s/login/%s/%s", conf.App.HttpProto, conf.App.Domain, adminAuth.AuthToken, initAdminUser.UUID)
	}
	bots := make([]dao.User, 0)
	err = store.Find(
		&bots,
		bolthold.
			Where("AuthLvl").
			Eq(enum.BOT).
			Index("AuthLvl"),
	)
	if err != nil {
		log.Println("Error when finding admins:", err)
	}
	if len(bots) == 0 {
		err := store.Insert(initBotUser.UUID.String(), initBotUser)
		if err != nil {
			log.Printf("%s alread exists in user table", initBotUser.Name)
		}
		conf.App.SetBotId(initBotUser.UUID)
	}
}
