package api

import (
	"fmt"
	"os"
	"testing"
	"time"

	data2 "github.com/anytimesoon/eurovision-party/pkg/data"
	dao2 "github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	service2 "github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

var (
	testDB *bolthold.Store

	adminUserId   = uuid.New()
	adminUserMock = dto2.User{
		UUID:      adminUserId,
		Name:      "admin",
		Slug:      "admin",
		Icon:      "default",
		AuthLvl:   authLvl.ADMIN,
		CreatedBy: adminUserId,
		CanInvite: true,
	}
	adminAuthMock = dto2.Auth{
		Token:      "adminToken",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     adminUserId,
		AuthLvl:    authLvl.ADMIN,
	}

	regularUserId   = uuid.New()
	regularUserMock = dto2.User{
		UUID:      regularUserId,
		Name:      "regular",
		Slug:      "regular",
		Icon:      "default",
		AuthLvl:   authLvl.USER,
		CreatedBy: adminUserId,
		CanInvite: true,
	}
	regularAuthMock = dto2.Auth{
		Token:      "regularToken",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     regularUserId,
		AuthLvl:    authLvl.USER,
	}
	countryNames = []string{"Austria", "Belgium", "Bulgaria", "Croatia", "Cyprus"}

	friendOfFriendUserId   = uuid.New()
	friendOfFriendUserMock = dao2.User{
		UUID:      friendOfFriendUserId,
		AuthLvl:   0,
		Name:      "friend of friend",
		Slug:      "friend-of-friend",
		Icon:      "default",
		Invites:   nil,
		CreatedBy: regularUserId,
		CanInvite: false,
	}
	freindOfFriendAuthMock = dto2.Auth{
		Token:      "fofToken",
		Expiration: time.Now().Add(time.Hour * 24 * 365),
		UserId:     friendOfFriendUserId,
		AuthLvl:    authLvl.FRIEND_OF_FRIEND,
	}

	testUserBroadcastChan = make(chan dto2.SocketMessage)
	testVoteBroadcastChan = make(chan dto2.SocketMessage)
)

func TestMain(m *testing.M) {
	testDB = setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() *bolthold.Store {
	var user dto2.User
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

func generateNewUsers() []*dto2.NewUser {
	newAdmin := &dto2.NewUser{
		Name:      adminUserMock.Name,
		Slug:      adminUserMock.Slug,
		UUID:      adminUserMock.UUID,
		AuthLvl:   adminUserMock.AuthLvl,
		Token:     adminAuthMock.Token,
		CreatedBy: adminUserId,
	}

	newRegular := &dto2.NewUser{
		Name:      regularUserMock.Name,
		Slug:      regularUserMock.Slug,
		UUID:      regularUserMock.UUID,
		AuthLvl:   regularUserMock.AuthLvl,
		Token:     regularAuthMock.Token,
		CreatedBy: adminUserId,
	}

	newFriendOfFriend := &dto2.NewUser{
		Name:      friendOfFriendUserMock.Name,
		Slug:      friendOfFriendUserMock.Slug,
		UUID:      friendOfFriendUserMock.UUID,
		AuthLvl:   friendOfFriendUserMock.AuthLvl,
		Token:     "",
		CreatedBy: regularUserId,
	}

	return []*dto2.NewUser{newAdmin, newRegular, newFriendOfFriend}
}

func newUsersFilteredById(id uuid.UUID) []*dto2.NewUser {
	filteredUsers := make([]*dto2.NewUser, 0)
	for _, user := range generateNewUsers() {
		if user.CreatedBy == id {
			filteredUsers = append(filteredUsers, user)
		}
	}
	return filteredUsers

}

func generateUsers(db *bolthold.Store) {
	err := db.Upsert(adminUserId.String(), dao2.User{}.FromDTO(adminUserMock))
	if err != nil {
		panic(err)
	}

	err = db.Upsert(adminAuthMock.Token, adminAuthMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(regularUserId.String(), dao2.User{}.FromDTO(regularUserMock))
	if err != nil {
		panic(err)
	}

	err = db.Upsert(regularAuthMock.Token, regularAuthMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(friendOfFriendUserId.String(), friendOfFriendUserMock)
	if err != nil {
		panic(err)
	}

	err = db.Upsert(freindOfFriendAuthMock.Token, freindOfFriendAuthMock)
	if err != nil {
		panic(err)
	}
}

func generateCountries(db *bolthold.Store) {

	for _, countryName := range countryNames {
		country := dto2.Country{
			Name: countryName,
			Slug: countryName,
		}
		err := db.Upsert(country.Slug, country)
		if err != nil {
			panic(err)
		}

		err = db.Upsert(country.Slug, dao2.VoteTracker{})
		if err != nil {
			panic(err)
		}
	}
}

func generateVotes(db *bolthold.Store) {
	countries := make([]dto2.Country, 0)
	err := db.Find(&countries, &bolthold.Query{})
	if err != nil {
		panic(err)
	}

	adminIdString := adminUserId.String()
	regularIdString := regularUserId.String()

	for i, countryName := range countryNames {
		err = db.Upsert(
			fmt.Sprintf("%s_%s", adminIdString, countryName),
			dao2.Vote{
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
			dao2.Vote{
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

	err = db.ReIndex(&dao2.Vote{}, []byte("Vote"))
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

func newTestAuthService() service2.AuthService {
	authRepository := data2.NewAuthRepositoryDB(testDB)
	sessionRepository := data2.NewSessionRepositoryDb(testDB)
	userRepositoryDb := data2.NewUserRepositoryDb(testDB)
	return service2.NewAuthService(authRepository, sessionRepository, userRepositoryDb, "5087c3b0928acd41f1907689a6f7bded8c42d03220934a7ad59e19c233a6bb7c")
}

func newTestVoteService() service2.VoteService {
	voteRepository := data2.NewVoteRepositoryDb(testDB)
	commentRepository := data2.NewCommentRepositoryDb(testDB)
	return service2.NewVoteService(voteRepository, testVoteBroadcastChan, commentRepository)
}

func newTestCountryService() service2.CountryService {
	countryRepository := data2.NewCountryRepositoryDb(testDB)
	return service2.NewCountryService(countryRepository)
}

func newTestUserService() service2.UserService {
	userRepository := data2.NewUserRepositoryDb(testDB)
	authRepository := data2.NewAuthRepositoryDB(testDB)
	commentRepository := data2.NewCommentRepositoryDb(testDB)
	voteRepository := data2.NewVoteRepositoryDb(testDB)
	return service2.NewUserService(
		userRepository,
		testUserBroadcastChan,
		authRepository,
		commentRepository,
		voteRepository,
	)
}

func newTestCommentService() service2.CommentService {
	commentRepository := data2.NewCommentRepositoryDb(testDB)
	return service2.NewCommentService(commentRepository, testUserBroadcastChan)
}

func newTestChatRoomService() *service2.Room {
	return service2.NewRoom()
}

func getAdminSession() string {
	authService := newTestAuthService()
	session, _ := authService.Login(adminAuthMock)
	return session.SessionToken
}

func getRegularSession() string {
	authService := newTestAuthService()
	session, _ := authService.Login(adminAuthMock)
	return session.SessionToken
}
