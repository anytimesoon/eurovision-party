package domain

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func NewUserRepositoryDb(db *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{db}
}

func (db UserRepositoryDb) CreateUser(userDTO dto.NewUser) (*NewUser, *errs.AppError) {
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

func (db UserRepositoryDb) VerifySlug(userDTO *dto.NewUser) error {
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

func (db UserRepositoryDb) UpdateUser(userDTO dto.User) (*User, *dto.Comment, *errs.AppError) {
	var user User

	updateUserQuery := "UPDATE user SET name = ? WHERE uuid = ?"
	getUserQuery := "SELECT * FROM user WHERE uuid = ?"
	addBotCommentQuery := "INSERT INTO comment(uuid, userId, text) VALUES (?, ?, ?)"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction to update user %s, %s", userDTO.UUID, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "user")
	}

	err = tx.Get(&user, getUserQuery, userDTO.UUID.String())
	if err != nil {
		log.Printf("Error while fetching user %s after update %s", userDTO.Name, err)
		return nil, nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	botComment := dto.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      fmt.Sprintf("🤖 %s changed their name to %s", user.Name, userDTO.Name),
		CreatedAt: time.Now(),
	}
	_, err = tx.Exec(addBotCommentQuery, botComment.UUID, botComment.UserId, botComment.Text)
	if err != nil {
		log.Printf("Error while writing bot comment after name update %s", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	_, err = tx.Exec(updateUserQuery, userDTO.Name, userDTO.UUID)
	if err != nil {
		log.Println("Error while updating user table", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "user")
	}

	err = tx.Get(&user, getUserQuery, userDTO.UUID.String())
	if err != nil {
		log.Printf("Error while fetching user %s after update %s", userDTO.Name, err)
		return nil, nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while committing user update for user %s. %s", userDTO.UUID, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "user")
	}

	return &user, &botComment, nil
}

func (db UserRepositoryDb) UpdateUserImage(id uuid.UUID) (*User, *dto.Comment, *errs.AppError) {
	var user User

	updateUserImageQuery := "UPDATE user SET icon = ? WHERE uuid = ?"
	getUserQuery := "SELECT * FROM user WHERE uuid = ?"
	addBotCommentQuery := "INSERT INTO comment(uuid, userId, text) VALUES (?, ?, ?)"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error while starting image transaction for user %s. %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	_, err = tx.Exec(updateUserImageQuery, id.String()+".png", id)
	if err != nil {
		log.Printf("Error while updating user image for user %s. %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	err = tx.Get(&user, getUserQuery, id.String())
	if err != nil {
		log.Printf("Error while fetching user %s after updating image %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	var botComment = dto.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      fmt.Sprintf("🤖 %s changed their picture", user.Name),
		CreatedAt: time.Now(),
	}
	_, err = tx.Exec(addBotCommentQuery, botComment.UUID, botComment.UserId, botComment.Text)
	if err != nil {
		log.Printf("Error while writing bot comment after updating image %s", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while committing image transaction for user %s. %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	return &user, &botComment, nil
}

func (db UserRepositoryDb) FindOneUser(slug string) (*User, *errs.AppError) {
	var user User

	query := "SELECT * FROM user WHERE slug = ?"
	err := db.client.Get(&user, query, slug)
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	return &user, nil
}

func (db UserRepositoryDb) DeleteUser(slug string) *errs.AppError {
	var user User

	query := "DELETE FROM user WHERE slug = ?"

	_, err := db.client.Exec(query, user, slug)
	if err != nil {
		log.Println("Error when deleting user", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "user")
	}

	return nil
}

func (db UserRepositoryDb) FindRegisteredUsers() (*[]NewUser, *errs.AppError) {
	users := make([]NewUser, 0)

	query := "SELECT u.uuid, u.name, u.slug, a.authToken FROM user u JOIN auth a ON u.uuid = a.userId WHERE u.authLvl NOT IN (3)"
	err := db.client.Select(&users, query)
	if err != nil {
		log.Println("Error while querying user table for registered users", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &users, nil
}
