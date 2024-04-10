package domain

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type AuthRepository interface {
	CreateUser(dto.NewUser) (*NewUser, *errs.AppError)
	Login(*dto.Auth) (*Auth, *User, *errs.AppError)
	Authorize(*dto.Auth) *errs.AppError
	AuthorizeChat(string, string) *errs.AppError
	VerifySlug(*dto.NewUser) error
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
	createSessionQuery := "UPDATE auth SET sessionToken = ? WHERE userId = ?"
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

func (db AuthRepositoryDB) Authorize(authDTO *dto.Auth) *errs.AppError {
	var auth Auth

	query := "SELECT * FROM auth WHERE sessionToken = ? and userId = ?"
	err := db.client.Get(&auth, query, authDTO.Token, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to authorize user %s and session token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	return nil
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

func (db AuthRepositoryDB) CreateUser(userDTO dto.NewUser) (*NewUser, *errs.AppError) {
	var user NewUser
	var auth Auth

	err := db.VerifySlug(&userDTO)
	if err != nil {
		log.Printf("Error when slufigying user %s with message %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	// Prepare queries for transaction
	newUserQuery := `INSERT INTO user(uuid, name, slug, authLvl) VALUES (?, ?, ?, 0)`
	newAuthQuery := `INSERT INTO auth(authToken, userId, authTokenExp, authLvl, lastUpdated, slug) VALUES (?, ?, NOW() + INTERVAL 10 DAY, 0, NOW(), ?)`
	findNewUserQuery := `SELECT u.uuid, u.name, u.slug, a.authToken FROM user u JOIN auth a ON u.uuid = a.userId WHERE u.uuid = ?`

	// Begin transaction that will create a new user and auth then return the new user
	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction for new auth for user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "auth")
	}

	userDTO.UUID = uuid.New()
	_, err = tx.Exec(newUserQuery, userDTO.UUID.String(), userDTO.Name, userDTO.Slug)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	auth.GenerateSecureToken(30)
	_, err = tx.Exec(newAuthQuery, auth.AuthToken, userDTO.UUID.String(), userDTO.Slug)
	if err != nil {
		log.Printf("Error when creating new auth for user %s, %s", userDTO.Name, err)
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "a new user")
	}

	err = tx.Get(&user, findNewUserQuery, userDTO.UUID)
	if err != nil {
		log.Printf("Error when retrieving new user %s after transaction. %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "a new user")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error when commiting auth transaction for new user %s. %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "a new user")
	}

	return &user, nil
}

func (db AuthRepositoryDB) VerifySlug(userDTO *dto.NewUser) error {
	// Verify the name is unique or add a number to the end
	counter := 0
	for {
		if counter > 0 {
			userDTO.Slug = userDTO.Slug + "-" + strconv.Itoa(counter)
		}

		query := fmt.Sprintf("SELECT * FROM user WHERE slug = '%s'", userDTO.Slug)
		rows, err := db.client.Query(query)
		if err != nil {
			return err
		}

		if !rows.Next() {
			break
		}

		counter++
	}

	return nil
}
