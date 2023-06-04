package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type AuthRepository interface {
	FindOneUserByEmail(string) (*User, *errs.AppError)
	CreateUser(dto.NewUser) (*NewUser, *errs.AppError)
	Login(*dto.Auth) (*Auth, *User, *errs.AppError)
	Authorize(*dto.Auth) (*Auth, *errs.AppError)
	AuthorizeChat(*dto.Auth) (*User, *errs.AppError)
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
	createSessionQuery := "UPDATE auth SET sessionToken = ?, sessionTokenExp = NOW() + INTERVAL 12 HOUR WHERE userId = ?"
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
		log.Printf("Unable to authorize user %s and etoken %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	return &auth, nil
}

func (db AuthRepositoryDB) AuthorizeChat(authDTO *dto.Auth) (*User, *errs.AppError) {
	var auth Auth

	getAuthQuery := "SELECT * FROM auth WHERE sessionToken = ? and userId = ?"
	err := db.client.Get(&auth, getAuthQuery, authDTO.Token, authDTO.UserId)
	if err != nil {
		log.Printf("Unable to authorize user %s and etoken %s combination for chat. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	var user User
	getUserQuery := "SELECT * FROM user WHERE uuid = ?"
	err = db.client.Get(&user, getUserQuery, authDTO.UserId.String())
	if err != nil {
		log.Printf("Unable to find user %s for chat. %s", authDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &user, nil
}

func (db AuthRepositoryDB) updateSession(userId uuid.UUID, auth *Auth) (*Auth, *errs.AppError) {
	updateAuthQuery := "UPDATE auth SET sessionToken = ?, sessionTokenExp = Now() + INTERVAL 1 DAY WHERE userId = ?"
	getAuthQuery := "SELECT * FROM auth WHERE userId = ?"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error while starting transaction to update auth for user %s. %s", userId, err)
		return auth, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}
	_, err = tx.Exec(updateAuthQuery, auth.SessionToken, userId)
	if err != nil {
		log.Printf("Error while updating auth for user %s. %s", userId, err)
		return auth, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	err = tx.Get(&auth, getAuthQuery, userId)
	if err != nil {
		log.Printf("Unable to find user %s when authorizing. %s", userId, err)
		return auth, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while committing transaction to update auth for user %s. %s", userId, err)
		return auth, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	return auth, nil
}

func (db AuthRepositoryDB) FindOneUserByEmail(email string) (*User, *errs.AppError) {
	var user User

	query := "SELECT * FROM user WHERE email = ?"
	err := db.client.Get(&user, query, email)
	if err != nil {
		log.Printf("Coudn't find a user with email address %s", email)
		return nil, errs.NewUnexpectedError("Coudn't find a user with email address" + email)
	}

	return &user, nil
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
	newUserQuery := `INSERT INTO user(uuid, name, email, slug, authLvl) VALUES (?, ?, ?, ?, 0)`
	newAuthQuery := `INSERT INTO auth(authToken, userId, authTokenExp, authLvl, slug) VALUES (?, ?, NOW() + INTERVAL 10 DAY, 0, ?)`
	findNewUserQuery := `SELECT u.uuid, u.name, u.email, u.slug, a.authToken FROM user u JOIN auth a ON u.uuid = a.userId WHERE u.uuid = ?`

	// Begin transaction that will create a new user and auth then return the new user
	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction for new auth for user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "auth")
	}

	userDTO.UUID = uuid.New()
	_, err = tx.Exec(newUserQuery, userDTO.UUID.String(), userDTO.Name, userDTO.Email, userDTO.Slug)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	auth.GenerateSecureToken(80)
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
