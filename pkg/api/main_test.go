package api

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/service"
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

	err = db.Upsert(adminUserId.String(), adminUserMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(adminAuthMock.Token, adminAuthMock)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Admin user created %v", user)

	return db
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
