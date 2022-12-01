package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindOneUserByEmail(string) *User
	CreateUser(dto.User) (*dto.Auth, *errs.AppError)
	Authenticate(*dto.Auth) (*Auth, *errs.AppError)
	Login(*dto.Auth) (*Auth, *errs.AppError)
}

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func NewAuthRepositoryDB(db *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{db}
}

func FindOneUserByEmail(email string) (*User, *errs.AppError) {
	return nil, nil
}

func (db AuthRepositoryDB) Authenticate(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE token = '%s' and userid = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authenticate user %s and token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't log you in. Please try again.")
	}

	// Auth found, so we can refresh
	var newAuth Auth
	newAuth.GenerateSecureToken()

	query = fmt.Sprintf(`UPDATE auth SET etoken = '%s', etexp = Now() + INTERVAL 1 DAY WHERE token = '%s'`, newAuth.Token, auth.Token)
	_, err = db.client.NamedExec(query, newAuth)
	if err != nil {
		log.Println("Error while updating auth", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	return &newAuth, nil
}

func (db AuthRepositoryDB) FindOneUserByEmail(email string) *User {
	var user User

	query := fmt.Sprintf(`SELECT * FROM user WHERE email = '%s'`, email)
	db.client.Get(&user, query)

	return &user
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

	auth.GenerateSecureToken()

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
