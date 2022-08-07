package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"eurovision/pkg/utils"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func NewUserRepositoryDb(db *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{db}
}

func (db UserRepositoryDb) FindAllUsers() ([]User, *errs.AppError) {
	users := make([]User, 0)

	query := "SELECT * FROM user"
	err := db.client.Select(&users, query)
	if err != nil {
		log.Println("Error while querying user table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return users, nil
}

func (db UserRepositoryDb) UpdateUser(userDTO dto.User) (*User, *errs.AppError) {
	var user User

	query := fmt.Sprintf(`UPDATE user SET name = '%s', icon = '%s' WHERE uuid = '%s'`, userDTO.Name, userDTO.Icon, userDTO.UUID.String())

	_, err := db.client.NamedExec(query, user)
	if err != nil {
		log.Println("Error while updating user table", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "user")
	}

	query = fmt.Sprintf(`SELECT * FROM user WHERE uuid = '%s'`, userDTO.UUID.String())
	err = db.client.Get(&user, query)
	if err != nil {
		log.Printf("Error while fetching user %s after update %s", userDTO.Name, err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	return &user, nil
}

func (db UserRepositoryDb) CreateUser(userDTO dto.User) (*User, *errs.AppError) {
	var user User

	slug := utils.Slugify(userDTO.Name)

	query := fmt.Sprintf(`INSERT INTO user(uuid, name, slug, authLvl) VALUES ('%s', '%s', '%s', 0)`, uuid.New().String(), userDTO.Name, slug)

	_, err := db.client.NamedExec(query, user)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	query = fmt.Sprintf(`SELECT * FROM user WHERE slug = '%s'`, slug)
	err = db.client.Get(&user, query)
	if err != nil {
		log.Printf("Error when fetching user %s after create %s", userDTO.Name, err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	return &user, nil
}

func (db UserRepositoryDb) FindOneUser(slug string) (*User, *errs.AppError) {
	var user User

	query := fmt.Sprintf(`SELECT * FROM user WHERE slug = '%s'`, slug)
	err := db.client.Get(&user, query)
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	return &user, nil
}

func (db UserRepositoryDb) DeleteUser(slug string) *errs.AppError {
	var user User

	query := fmt.Sprintf(`DELETE FROM user WHERE slug = '%s'`, slug)

	_, err := db.client.NamedExec(query, user)
	if err != nil {
		log.Println("Error when deleting user", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "user")
	}

	return nil
}
