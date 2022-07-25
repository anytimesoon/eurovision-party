package domain

import (
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
