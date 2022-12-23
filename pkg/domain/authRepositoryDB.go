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
	CreateUser(dto.NewUser) (*Auth, *errs.AppError)
	Login(*dto.Auth) (*Auth, *errs.AppError)
	Authorize(*dto.Auth) (*Auth, *errs.AppError)
	VerifySlug(*dto.NewUser) error
}

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func NewAuthRepositoryDB(db *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{db}
}

func (db AuthRepositoryDB) Login(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE token = '%s' and userId = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authenticate user %s and token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't log you in. Please try again.")
	}

	// Auth found, so we can refresh
	auth.GenerateSecureEToken(20)

	query = fmt.Sprintf(`UPDATE auth SET etoken = '%s', etexp = Now() + INTERVAL 1 DAY WHERE userId = '%s'`, auth.EToken, auth.UserId)
	_, err = db.client.NamedExec(query, auth)
	if err != nil {
		log.Println("Error while updating auth", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	query = fmt.Sprintf(`SELECT * FROM auth WHERE userId = '%s'`, authDTO.UserId)
	err = db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to find user %s. %s", authDTO.UserId, err)
		return nil, errs.NewUnauthorizedError("Couldn't log you in. Please try again.")
	}

	return &auth, nil
}

func (db AuthRepositoryDB) Authorize(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE etoken = '%s' and userId = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authorize user %s and etoken %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't log you in. Please try again.")
	}

	// TODO this whole section needs looking at. It would be better to update time on DAO and then send to DB.
	// Auth found, so we can refresh
	auth.GenerateSecureEToken(20)

	query = fmt.Sprintf(`UPDATE auth SET etoken = '%s', etexp = Now() + INTERVAL 1 DAY WHERE userId = '%s'`, auth.EToken, auth.UserId)
	_, err = db.client.NamedExec(query, auth)
	if err != nil {
		log.Println("Error while updating auth", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	query = fmt.Sprintf(`SELECT * FROM auth WHERE userId = '%s'`, authDTO.UserId)
	err = db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to find user %s. %s", authDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	return &auth, nil
}

func (db AuthRepositoryDB) FindOneUserByEmail(email string) (*User, *errs.AppError) {
	var user User

	query := fmt.Sprintf(`SELECT * FROM user WHERE email = '%s'`, email)
	err := db.client.Get(&user, query)
	if err != nil {
		log.Printf("Coudn't find a user with email address %s", email)
		return nil, errs.NewUnexpectedError("Coudn't find a user with email address" + email)
	}

	return &user, nil
}

func (db AuthRepositoryDB) CreateUser(userDTO dto.NewUser) (*Auth, *errs.AppError) {
	var user User
	var auth Auth

	err := db.VerifySlug(&userDTO)
	if err != nil {
		log.Printf("Error when slufigying user %s with message %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	user.UUID = uuid.New()

	query := fmt.Sprintf(`INSERT INTO user(uuid, name, email, slug, authLvl) VALUES ('%s', '%s', '%s', '%s', 0)`, user.UUID.String(), userDTO.Name, userDTO.Email, userDTO.Slug)

	_, err = db.client.NamedExec(query, user)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	auth.GenerateSecureToken(80)

	query = fmt.Sprintf(`INSERT INTO auth(token, userId, expiration, authLvl, slug) VALUES ('%s', '%s', NOW() + INTERVAL 10 DAY, %d, '%s')`, auth.Token, user.UUID, user.AuthLvl, user.Slug)
	_, err = db.client.NamedExec(query, auth)
	if err != nil {
		log.Printf("Error when creating new auth for user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "auth")
	}

	return &auth, nil
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
