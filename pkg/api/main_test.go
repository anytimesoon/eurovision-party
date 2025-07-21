package api

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
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
		Token:      "adminToken",
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
		Token:      "regularToken",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     regularUserId,
		AuthLvl:    0,
	}
	countryNames = []string{"Austria", "Belgium", "Bulgaria", "Croatia", "Cyprus"}

	testUserBroadcastChan = make(chan dto.SocketMessage)
	testVoteBroadcastChan = make(chan dto.SocketMessage)
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

func generateNewUsers() []*dto.NewUser {
	newAdmin := &dto.NewUser{
		Name:    adminUserMock.Name,
		Slug:    adminUserMock.Slug,
		UUID:    adminUserMock.UUID,
		AuthLvl: adminUserMock.AuthLvl,
		Token:   adminAuthMock.Token,
	}

	newRegular := &dto.NewUser{
		Name:    regularUserMock.Name,
		Slug:    regularUserMock.Slug,
		UUID:    regularUserMock.UUID,
		AuthLvl: regularUserMock.AuthLvl,
		Token:   regularAuthMock.Token,
	}

	return []*dto.NewUser{newAdmin, newRegular}
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

		err = db.Upsert(country.Slug, dao.VoteTracker{})
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

func newTestAuthService() service.AuthService {
	authRepository := data.NewAuthRepositoryDB(testDB)
	sessionRepository := data.NewSessionRepositoryDb(testDB)
	userRepositoryDb := data.NewUserRepositoryDb(testDB)
	return service.NewAuthService(authRepository, sessionRepository, userRepositoryDb, "5087c3b0928acd41f1907689a6f7bded8c42d03220934a7ad59e19c233a6bb7c")
}

func newTestVoteService() service.VoteService {
	voteRepository := data.NewVoteRepositoryDb(testDB)
	commentRepository := data.NewCommentRepositoryDb(testDB)
	return service.NewVoteService(voteRepository, testVoteBroadcastChan, commentRepository)
}

func newTestCountryService() service.CountryService {
	countryRepository := data.NewCountryRepositoryDb(testDB)
	return service.NewCountryService(countryRepository)
}

func newTestUserService() service.UserService {
	userRepository := data.NewUserRepositoryDb(testDB)
	authRepository := data.NewAuthRepositoryDB(testDB)
	commentRepository := data.NewCommentRepositoryDb(testDB)
	voteRepository := data.NewVoteRepositoryDb(testDB)
	return service.NewUserService(
		userRepository,
		testUserBroadcastChan,
		authRepository,
		commentRepository,
		voteRepository,
	)
}

func newTestCommentService() service.CommentService {
	commentRepository := data.NewCommentRepositoryDb(testDB)
	return service.NewCommentService(commentRepository)
}

func newTestChatRoomService() *service.Room {
	return service.NewRoom(newTestCommentService())
}

func getAdminSession() string {
	authService := newTestAuthService()
	auth, _, _ := authService.Login(adminAuthMock)
	return auth.Token
}

func getRegularSession() string {
	authService := newTestAuthService()
	auth, _, _ := authService.Login(regularAuthMock)
	return auth.Token
}
