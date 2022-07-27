package domain

import (
	"eurovision/pkg/dto"
	"fmt"
	"log"

	"github.com/google/uuid"
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

func (db CommentRepositoryDb) CreateComment(commentDTO dto.Comment) (Comment, error) {
	var comment Comment

	uuid := uuid.New().String()

	query := fmt.Sprintf(`INSERT INTO comment(uuid, userId, text) VALUES ('%s', '%s', '%s')`, uuid, commentDTO.UserId, commentDTO.Text)

	_, err := db.client.NamedExec(query, comment)
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", commentDTO.UserId, err)
		return comment, err
	}

	query = fmt.Sprintf(`SELECT * FROM comment WHERE uuid = '%s'`, uuid)
	err = db.client.Get(&comment, query)
	if err != nil {
		log.Printf("Error when fetching comment after create %s", err)
		return comment, err
	}

	return comment, nil
}
