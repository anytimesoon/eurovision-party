package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthRepository interface {
	Login(*dto.Auth) (*Auth, *User, *errs.AppError)
	Authorize(*dto.Auth) (*Auth, *errs.AppError)
	AuthorizeChat(string, string) *errs.AppError
}

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func NewAuthRepositoryDB(db *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{db}
}

func (db AuthRepositoryDB) Login(authDTO *dto.Auth) (*Auth, *User, *errs.AppError) {
	var auth Auth
	var user User

	getAuthQuery := "SELECT * FROM auth WHERE authToken = ? and userId = ?"
	createSessionQuery := "UPDATE auth SET sessionToken = ?, sessionTokenExp = NOW() + INTERVAL 7 DAY WHERE userId = ?"
	getUserQuery := "SELECT * FROM user WHERE uuid = ?"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction login user %s, %s", authDTO.UserId, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	err = tx.Get(&auth, getAuthQuery, authDTO.Token, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to authenticate user %s and token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	auth.GenerateSecureSessionToken(20)
	_, err = tx.Exec(createSessionQuery, auth.SessionToken, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to generate new session token for user %s. %s", authDTO.UserId, err)
		return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	err = tx.Get(&user, getUserQuery, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to find user %s. %s", authDTO.UserId, err)
		return nil, nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error when commiting login transaction for user %s. %s", authDTO.UserId, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	return &auth, &user, nil
}

func (db AuthRepositoryDB) Authorize(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := "SELECT * FROM auth WHERE sessionToken = ? and userId = ?"
	err := db.client.Get(&auth, query, authDTO.Token, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to authorize user %s and session token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	return &auth, nil
}

func (db AuthRepositoryDB) AuthorizeChat(token, userId string) *errs.AppError {
	var auth Auth

	getAuthQuery := "SELECT * FROM auth WHERE authToken = ? AND userId = ?"
	err := db.client.Get(&auth, getAuthQuery, token, userId)
	if err != nil {
		log.Printf("Unable to find user %s for chat. %s", userId, err)
		return errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return nil
}
