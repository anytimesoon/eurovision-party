package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindOneUserByEmail(string) *User
	CreateUser(dto.User) (*dto.Auth, *errs.AppError)
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

func (db AuthRepositoryDB) FindOneUserByEmail(email string) *User {
	var user User

	query := fmt.Sprintf(`SELECT * FROM user WHERE email = '%s'`, email)
	db.client.Get(&user, query)

	return &user
}

func (db AuthRepositoryDB) CreateUser(userDTO dto.NewUser, authDTO dto.Auth) (*dto.Auth, *errs.AppError) {
	var user User

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

	authDTO.UserId = user.UUID
	authDTO.Expiration = time.Now().Add(time.Hour * (24 * 7))

	return &authDTO, nil
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
