package api

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
	"os"
	"testing"
	"time"
)

var (
	testDB *bolthold.Store

	adminUserId   = uuid.New()
	adminUserMock = dto.User{
		UUID:    adminUserId,
		Name:    "admin",
		Slug:    "admin",
		Icon:    "default",
		AuthLvl: 1,
	}
	adminAuthMock = dto.Auth{
		Token:      "token",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     adminUserId,
		AuthLvl:    1,
	}

	regularUserId   = uuid.New()
	regularUserMock = dto.User{
		UUID:    regularUserId,
		Name:    "regular",
		Slug:    "regular",
		Icon:    "default",
		AuthLvl: 0,
	}
	regularAuthMock = dto.Auth{
		Token:      "token",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     regularUserId,
		AuthLvl:    0,
	}
	countryNames = []string{"Austria", "Belgium", "Bulgaria", "Croatia", "Cyprus"}
)

func TestMain(m *testing.M) {
	testDB = setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() *bolthold.Store {
	var user dto.User
	db, err := bolthold.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	generateUsers(db)
	generateCountries(db)
	generateVotes(db)

	fmt.Printf("Admin user created %v", user)

	return db
}

func generateUsers(db *bolthold.Store) {
	err := db.Upsert(adminUserId.String(), adminUserMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(adminAuthMock.Token, adminAuthMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(regularUserId.String(), regularUserMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(regularAuthMock.Token, regularAuthMock)
	if err != nil {
		panic(err)
	}
}

func generateCountries(db *bolthold.Store) {

	for _, countryName := range countryNames {
		country := dto.Country{
			Name: countryName,
			Slug: countryName,
		}
		err := db.Upsert(country.Slug, country)
		if err != nil {
			panic(err)
		}
	}
}

func generateVotes(db *bolthold.Store) {
	countries := make([]dto.Country, 0)
	err := db.Find(&countries, &bolthold.Query{})
	if err != nil {
		panic(err)
	}

	adminIdString := adminUserId.String()
	regularIdString := regularUserId.String()

	for i, countryName := range countryNames {
		err = db.Upsert(
			fmt.Sprintf("%s_%s", adminIdString, countryName),
			dao.Vote{
				UserId:      adminUserId,
				CountrySlug: countryName,
				Costume:     uint8(i),
				Song:        uint8(i),
				Performance: uint8(i),
				Props:       uint8(i),
				Total:       i * 4,
			})
		if err != nil {
			panic(err)
		}

		err = db.Upsert(
			fmt.Sprintf("%s_%s", regularIdString, countryName),
			dao.Vote{
				UserId:      regularUserId,
				CountrySlug: countryName,
				Costume:     uint8(i),
				Song:        uint8(i),
				Performance: uint8(i),
				Props:       uint8(i),
				Total:       i * 4,
			})
		if err != nil {
			panic(err)
		}
	}

	err = db.ReIndex(&dao.Vote{}, []byte("Vote"))
}

func shutdown() {
	err := testDB.Close()
	if err != nil {
		panic(err)
	}
	err = os.Remove("test.db")
	if err != nil {
		panic(err)
	}
}

func newMockAuthService() service.AuthService {
	authRepository := data.NewAuthRepositoryDB(testDB)
	sessionRepository := data.NewSessionRepositoryDb(testDB)
	userRepositoryDb := data.NewUserRepositoryDb(testDB)
	return service.NewAuthService(authRepository, sessionRepository, userRepositoryDb)
}

func newMockVoteService() service.VoteService {
	voteRepository := data.NewVoteRepositoryDb(testDB)
	return service.NewVoteService(voteRepository)
}
