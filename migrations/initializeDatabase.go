package migrations

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/timshannon/bolthold"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type resultCount []uint8

func Start(store *bolthold.Store) sqlx.DB {
	sqlDb := sqlx.MustConnect("mysql", dsn())
	log.Println("Successfully connected to database")

	log.Println("Building tables 🏗")
	//CreateAuthTable(sqlDb)
	//hasCountries := CreateCountriesTable(sqlDb)
	//hasUsers := CreateUsersTable(sqlDb)
	//CreateCommentsTable(sqlDb)
	CreateVotesTable(sqlDb)

	log.Println("Seeding tables 🌱")
	//if !hasCountries {
	//	AddCountries(sqlDb)
	//}
	//
	//if !hasUsers {
	//	AddUsers(sqlDb)
	//}

	addCountriesBolt(store)
	addUsersBolt(store)

	return *sqlDb
}

func addUsersBolt(store *bolthold.Store) {
	admins := make([]domain.User, 0)
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
		err := store.Insert(initAdminUser.UUID.String(), initAdminUser)
		if err != nil {
			log.Printf("%s alread exists in user table", initAdminUser.Name)
		}

		adminAuth := domain.Auth{
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

		log.Printf("%s%s/login/%s/%s", conf.App.HttpProto, conf.App.Domain, adminAuth.AuthToken, initAdminUser.UUID)
	}
	bots := make([]domain.User, 0)
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

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.App.DbUsername, conf.App.DbPassword, conf.App.DbHostname, conf.App.DbPort, conf.App.DbName)
}

func addCountriesBolt(store *bolthold.Store) {
	for _, country := range initCountriesWithParticipating {
		err := store.Insert(country.Slug, country)
		if err != nil {
			log.Printf("Skipping %s %s: already exists in country table", country.Flag, country.Name)
		}
	}
}
