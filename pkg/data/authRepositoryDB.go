package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/timshannon/bolthold"
	"log"
)

type AuthRepository interface {
	GetAuth(string) (*dao.Auth, *errs.AppError)
	UpdateAuth(*dao.Auth) error
}

type AuthRepositoryDB struct {
	store *bolthold.Store
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
