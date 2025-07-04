package data

import (
	"errors"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
	"log"
	"strconv"
)

type UserRepository interface {
	CreateUser(dao.User) (*dao.User, error)
	GetAllUsers() ([]dao.User, error)
	GetOneUserBySlug(string) (*dao.User, error)
	GetOneUserById(uuid.UUID) (*dao.User, error)
	DeleteUser(string) error
	GetRegisteredUsers() (*[]dao.User, error)
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

	err := db.verifySlug(&user)
	if err != nil {
		log.Printf("Error when slufigying user %s with message %s", user.Name, err)
		return nil, err
	}

	err = db.store.Insert(user.UUID.String(), user)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", user.Name, err)
		return nil, err
	}

	return &user, nil
}

func (db UserRepositoryDb) verifySlug(newUser *dao.User) error {
	// Verify the name is unique or add a number to the end
	counter := 0
	for {
		var user dao.User
		err := db.store.FindOne(&user, bolthold.Where("Slug").Eq(newUser.Slug))
		if err != nil {
			if errors.Is(err, bolthold.ErrNotFound) {
				// no users with this slug found, so validation is complete
				if counter > 0 {
					newUser.Slug = newUser.Slug[:len(newUser.Slug)-counter] + "-" + strconv.Itoa(counter)
				}
				return nil
			}
			return err
		}

		counter++
		newUser.Slug = newUser.Slug + "i"
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

func (db UserRepositoryDb) DeleteUser(slug string) error {
	user, err := db.GetOneUserBySlug(slug)
	if err != nil {
		return err
	}

	err = db.store.Delete(user.UUID.String(), user)
	if err != nil {
		log.Println("Error when deleting user", err)
		return err
	}

	return nil
}

func (db UserRepositoryDb) GetRegisteredUsers() (*[]dao.User, error) {
	users := make([]dao.User, 0)

	err := db.store.Find(&users, bolthold.Where("AuthLvl").Ne(enum.BOT))
	if err != nil {
		log.Println("Error while querying user table for registered users", err)
		return nil, err
	}

	return &users, nil
}
