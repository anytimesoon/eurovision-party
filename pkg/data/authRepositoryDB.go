package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/timshannon/bolthold"
	"log"
	"time"
)

type AuthRepository interface {
	Login(*dto.Auth) (*dao.Auth, *dao.User, *errs.AppError)
	Authorize(*dto.Auth) (*dao.Auth, *errs.AppError)
}

type AuthRepositoryDB struct {
	store *bolthold.Store
}

func NewAuthRepositoryDB(store *bolthold.Store) AuthRepositoryDB {
	return AuthRepositoryDB{store}
}

func (db AuthRepositoryDB) Login(authDTO *dto.Auth) (*dao.Auth, *dao.User, *errs.AppError) {
	var auth dao.Auth
	var user dao.User

	err := db.store.Get(authDTO.Token, &auth)
	if err != nil {
		log.Printf("Unable to authenticate user %s and token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	auth.GenerateSecureSessionToken(20)
	auth.SessionTokenExp = time.Now().Add(7 * 24 * time.Hour)
	err = db.store.Update(auth.AuthToken, auth)
	if err != nil {
		log.Printf("Unable to generate new session token for user %s. %s", authDTO.UserId, err)
		return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	err = db.store.Upsert(auth.SessionToken,
		dao.Session{
			AuthToken:    auth.AuthToken,
			SessionToken: auth.SessionToken,
			UserId:       authDTO.UserId},
	)

	err = db.store.Get(authDTO.UserId, &user)
	if err != nil {
		log.Printf("Unable to find user during login, trying again. %s", err)
		err = db.store.FindOne(&user, bolthold.Where("UUID").Eq(authDTO.UserId))
		if err != nil {
			log.Printf("Unable to find user %s during login. %s", authDTO.UserId, err)
			return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
		}
		//err = db.store.Update(user.UUID, user)
		users := make([]dao.User, 0)
		err = db.store.Find(&users, bolthold.Where(bolthold.Key).Eq(authDTO.UserId))
		if err != nil {
			log.Printf("What the fuck, dude")
		}
	}

	return &auth, &user, nil
}

func (db AuthRepositoryDB) Authorize(authDTO *dto.Auth) (*dao.Auth, *errs.AppError) {
	var auth dao.Auth
	var user dao.User
	var session dao.Session

	err := db.store.Get(authDTO.Token, &session)
	if err != nil {
		log.Println("Unable to find session", err)
		return nil, errs.NewUnauthorizedError(errs.Common.Login)
	}
	//err = db.store.FindOne(&user, bolthold.Where(bolthold.Key).Eq(authDTO.UserId))
	err = db.store.Get(authDTO.UserId.String(), &user)
	if err != nil {
		log.Println("Unable to find user during auth", err)
		return nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	err = db.store.Get(session.AuthToken, &auth)
	if err != nil {
		log.Println("Unable to find auth", err)
		return nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	return &auth, nil
}
