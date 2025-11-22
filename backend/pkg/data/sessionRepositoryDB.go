package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
)

type SessionRepository interface {
	UpsertSession(string, string, uuid.UUID) error
	GetSession(string) (*dao.Session, error)
	DeleteSession(string) error
}

type SessionRepositoryDB struct {
	store *bolthold.Store
}

func NewSessionRepositoryDb(store *bolthold.Store) SessionRepositoryDB {
	return SessionRepositoryDB{store}
}

func (db SessionRepositoryDB) UpsertSession(authToken string, sessionToken string, userId uuid.UUID) error {
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

func (db SessionRepositoryDB) DeleteSession(sessionToken string) error {
	return db.store.Delete(sessionToken, dao.Session{})
}
