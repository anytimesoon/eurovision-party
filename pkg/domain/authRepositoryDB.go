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
	Authenticate(string, string) (*Auth, *errs.AppError)
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

func (db AuthRepositoryDB) Authenticate(token string, userId string) (*Auth, *errs.AppError) {
	var auth Auth

	query := fmt.Sprintf(`SELECT * FROM auth WHERE token = '%s'`, token)
	db.client.Get(&auth, query)

	return &auth, nil
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

	query = fmt.Sprintf(`INSERT INTO auth(token, userId, expiration) VALUES ('%s', '%s', DATE_ADD(NOW(), INTERVAL 5 DAY))`, auth.Token, user.UUID)
	_, err = db.client.NamedExec(query, user)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	query = fmt.Sprintf(`SELECT * FROM auth WHERE token = '%s'`, auth.Token)
	err = db.client.Get(&auth, query)
	if err != nil {
		log.Println("Error when fetching auth after create:", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your authentication token")
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
