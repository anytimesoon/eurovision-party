package data

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

type UserRepository interface {
	CreateUser(dao.User) (*dao.User, error)
	GetAllUsers() ([]dao.User, error)
	GetOneUserBySlug(string) (*dao.User, error)
	GetOneUserById(uuid.UUID) (*dao.User, error)
	GetRegisteredUsers(dao.User) (*[]dao.User, error)
	UpdateUser(dao.User) (*dao.User, error)
	UpdateUserImage(uuid.UUID) (*dao.User, error)
}

type UserRepositoryDb struct {
	store *bolthold.Store
}

func NewUserRepositoryDb(store *bolthold.Store) UserRepositoryDb {
	return UserRepositoryDb{store}
}

func (db UserRepositoryDb) GetOneUserBySlug(slug string) (*dao.User, error) {
	var user dao.User
	err := db.store.FindOne(&user, bolthold.Where("Slug").Eq(slug))
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, err
	}
	return &user, nil
}

func (db UserRepositoryDb) GetOneUserById(userId uuid.UUID) (*dao.User, error) {
	var user dao.User
	err := db.store.Get(userId.String(), &user)
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, err
	}
	return &user, nil
}

func (db UserRepositoryDb) CreateUser(user dao.User) (*dao.User, error) {

	users := make([]dao.User, 0)
	slugifiedUser, err := db.verifySlug(&user)
	if err != nil {
		log.Printf("Error when slufigying user %s with message %s", user.Name, err)
		return nil, err
	}

	err = db.store.Find(&users, &bolthold.Query{})
	matches := make([]dao.User, 0)
	for _, user := range users {
		if strings.Contains(user.Slug, "asdf") {
			matches = append(matches, user)
		}
	}

	err = db.store.Insert(user.UUID.String(), slugifiedUser)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", user.Name, err)
		return nil, err
	}

	return slugifiedUser, nil
}

func (db UserRepositoryDb) verifySlug(newUser *dao.User) (*dao.User, error) {
	// Verify the name is unique or add a number to the end
	originalSlug := newUser.Slug
	counter := 0
	for {
		var user dao.User
		err := db.store.FindOne(&user, bolthold.Where("Slug").Eq(newUser.Slug))
		if err != nil {
			if errors.Is(err, bolthold.ErrNotFound) {
				return newUser, nil
			}
			return nil, err
		}

		counter++
		newUser.Slug = originalSlug + "-" + strconv.Itoa(counter)
		if counter > 100 {
			return nil, errors.New("unable to generate unique slug after 100 attempts")
		}
	}
}

func (db UserRepositoryDb) GetAllUsers() ([]dao.User, error) {
	users := make([]dao.User, 0)

	err := db.store.Find(&users, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying user table", err)
		return nil, err
	}

	return users, nil
}

func (db UserRepositoryDb) UpdateUser(user dao.User) (*dao.User, error) {
	err := db.store.Update(user.UUID.String(), user)
	if err != nil {
		log.Println("Error while updating user table", err)
		return nil, err
	}

	return &user, nil
}

func (db UserRepositoryDb) UpdateUserImage(id uuid.UUID) (*dao.User, error) {
	var user dao.User

	err := db.store.Get(id.String(), &user)
	if err != nil {
		log.Printf("Error while fetching user %s to image %s", id, err)
		return nil, err
	}

	user.Icon = id.String() + ".png"
	err = db.store.Update(id.String(), user)
	if err != nil {
		log.Printf("Error while updating user image for user %s. %s", id, err)
		return nil, err
	}

	return &user, nil
}

func (db UserRepositoryDb) GetRegisteredUsers(user dao.User) (*[]dao.User, error) {
	users := make([]dao.User, 0)
	var err error

	if user.AuthLvl == authLvl.ADMIN {
		err = db.store.Find(&users, bolthold.Where("AuthLvl").Ne(authLvl.BOT).And("IsBanned").Eq(false).SortBy("Name"))
	} else {
		err = db.store.Find(&users, bolthold.Where("CreatedBy").Eq(user.UUID).And("IsBanned").Eq(false).SortBy("Name"))
	}
	if err != nil {
		log.Println("Error while querying user table for registered users", err)
		return nil, err
	}

	return &users, nil
}
