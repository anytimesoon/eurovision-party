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
	Login(*dto.Auth) (*Auth, *errs.AppError)
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

func (db AuthRepositoryDB) Login(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE token = '%s' and userId = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authenticate user %s and token %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't log you in. Please try again.")
	}

	// Auth found, so we can renew
	newAuth, appErr := db.updateNewToken(authDTO.UserId)
	if appErr != nil {
		return nil, appErr
	}

	return newAuth, nil
}

func (db AuthRepositoryDB) Authorize(authDTO *dto.Auth) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE etoken = '%s' and userId = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authorize user %s and etoken %s combination. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	// Auth found, so we can renew
	newAuth, appErr := db.updateNewToken(authDTO.UserId)
	if appErr != nil {
		return nil, appErr
	}

	return newAuth, nil
}

func (db AuthRepositoryDB) AuthorizeChat(authDTO *dto.Auth) (*User, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE etoken = '%s' and userId = '%s'`, authDTO.Token, authDTO.UserId)
	err := db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to authorize user %s and etoken %s combination for chat. %s", authDTO.UserId, authDTO.Token, err)
		return nil, errs.NewUnauthorizedError("Couldn't authenticate you. Please try again.")
	}

	var user User
	query = fmt.Sprintf(`SELECT * FROM user WHERE uuid = '%s'`, authDTO.UserId.String())
	err = db.client.Get(&user, query)
	if err != nil {
		log.Printf("Unable to find user %s for chat. %s", authDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &user, nil
}

func (db AuthRepositoryDB) updateNewToken(userId uuid.UUID) (*Auth, *errs.AppError) {
	var auth Auth
	auth.GenerateSecureEToken(20)

	query := fmt.Sprintf(`UPDATE auth SET etoken = '%s', etexp = Now() + INTERVAL 1 DAY WHERE userId = '%s'`, auth.EToken, userId)
	_, err := db.client.Exec(query)
	if err != nil {
		log.Println("Error while updating auth", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your authentication")
	}

	query = fmt.Sprintf(`SELECT * FROM auth WHERE userId = '%s'`, userId)
	err = db.client.Get(&auth, query)
	if err != nil {
		log.Printf("Unable to find user %s. %s", userId, err)
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
	newAuthQuery := `INSERT INTO auth(token, userId, texp, authLvl, slug) VALUES (?, ?, NOW() + INTERVAL 10 DAY, 0, ?)`
	findNewUserQuery := `SELECT u.uuid, u.name, u.email, u.slug, a.token FROM user u JOIN auth a ON u.uuid = a.userId WHERE u.uuid = ?`

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
	_, err = tx.Exec(newAuthQuery, auth.Token, userDTO.UUID.String(), userDTO.Slug)
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
