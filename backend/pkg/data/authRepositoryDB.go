package data

import (
	"log"

	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

type AuthRepository interface {
	GetAuth(string) (*dao.Auth, *errs.AppError)
	UpdateAuth(*dao.Auth) error
	CreateAuth(dao.Auth) (*dao.Auth, error)
	GetAuthFromUserId(uuid uuid.UUID) (*dao.Auth, error)
}

type AuthRepositoryDB struct {
	store *bolthold.Store
}

func (db AuthRepositoryDB) GetAuthFromUserId(userId uuid.UUID) (*dao.Auth, error) {
	var auth dao.Auth
	err := db.store.FindOne(&auth, bolthold.Where("UserId").Eq(userId))
	if err != nil {
		log.Println("Error while querying auth table for registered user.", err)
		return nil, err
	}
	return &auth, nil
}

func (db AuthRepositoryDB) CreateAuth(auth dao.Auth) (*dao.Auth, error) {
	err := db.store.Insert(auth.AuthToken, auth)
	if err != nil {
		log.Printf("Error when creating new auth for user %s, %s", auth.UserId, err)
		return nil, err
	}
	return &auth, nil
}

func NewAuthRepositoryDB(store *bolthold.Store) AuthRepositoryDB {
	return AuthRepositoryDB{store}
}

func (db AuthRepositoryDB) GetAuth(authToken string) (*dao.Auth, *errs.AppError) {
	var auth dao.Auth
	err := db.store.Get(authToken, &auth)
	if err != nil {
		log.Println("Unable to find auth", err)
		return nil, errs.NewUnauthorizedError(errs.Common.Login)
	}
	return &auth, nil
}

func (db AuthRepositoryDB) UpdateAuth(auth *dao.Auth) error {
	err := db.store.Update(auth.AuthToken, auth)
	if err != nil {
		log.Println("Unable to update auth", err)
		return err
	}
	return nil
}
