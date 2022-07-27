package domain

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

type CommentRepositoryDb struct {
	client *sqlx.DB
}

func NewCommentRepositoryDb(db *sqlx.DB) CommentRepositoryDb {
	return CommentRepositoryDb{db}
}

func (db CommentRepositoryDb) FindAllComments() ([]Comment, error) {
	comments := make([]Comment, 0)

	query := "SELECT * FROM country"
	err := db.client.Select(&comments, query)
	if err != nil {
		log.Println("Error while querying country table", err)
		return nil, err
	}

	return comments, nil
}
