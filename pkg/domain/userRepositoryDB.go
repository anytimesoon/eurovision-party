package domain

import (
	"eurovision/pkg/dto"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func NewUserRepositoryDb(db *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{db}
}

func (db UserRepositoryDb) FindAllUsers() ([]User, error) {
	users := make([]User, 0)

	query := "SELECT * FROM user"
	err := db.client.Select(&users, query)
	if err != nil {
		log.Println("Error while querying user table", err)
		return nil, err
	}

	return users, nil
}

func (db UserRepositoryDb) UpdateUser(userDTO dto.User) (User, error) {
	var user User

	query := fmt.Sprintf(`UPDATE user SET name = '%s', icon = '%s' WHERE uuid = '%s'`, userDTO.Name, userDTO.Icon, userDTO.UUID.String())

	_, err := db.client.NamedExec(query, user)
	if err != nil {
		log.Println("Error while updating country table", err)
		return user, err
	}

	query = fmt.Sprintf(`SELECT * FROM user WHERE uuid = '%s'`, userDTO.UUID.String())
	err = db.client.Get(&user, query)
	if err != nil {
		log.Printf("Error while fetching user %s after update %s", userDTO.Name, err)
		return user, err
	}

	return user, nil
}
