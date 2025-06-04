package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

type SessionRepository interface {
	UpdateSession(string, string, uuid.UUID) error
	GetSession(string) (*dao.Session, error)
}

type SessionRepositoryDB struct {
	store *bolthold.Store
}

func NewSessionRepositoryDb(store *bolthold.Store) SessionRepositoryDB {
	return SessionRepositoryDB{store}
}

func (db SessionRepositoryDB) UpdateSession(authToken string, sessionToken string, userId uuid.UUID) error {
	err := db.store.Upsert(sessionToken, dao.Session{
		AuthToken:    authToken,
		SessionToken: sessionToken,
		UserId:       userId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (db SessionRepositoryDB) GetSession(sessionToken string) (*dao.Session, error) {
	var session dao.Session
	err := db.store.Get(sessionToken, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
