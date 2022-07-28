package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
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

func (db CommentRepositoryDb) FindAllComments() ([]Comment, *errs.AppError) {
	comments := make([]Comment, 0)

	query := "SELECT * FROM comment"
	err := db.client.Select(&comments, query)
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	return comments, nil
}

func (db CommentRepositoryDb) CreateComment(commentDTO dto.Comment) (*Comment, *errs.AppError) {
	var comment Comment

	uuid := uuid.New().String()

	query := fmt.Sprintf(`INSERT INTO comment(uuid, userId, text) VALUES ('%s', '%s', '%s')`, uuid, commentDTO.UserId, commentDTO.Text)

	_, err := db.client.NamedExec(query, comment)
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", commentDTO.UserId, err)
		return nil, errs.NewUnexpectedError("Something went wrong when adding your comment")
	}

	query = fmt.Sprintf(`SELECT * FROM comment WHERE uuid = '%s'`, uuid)
	err = db.client.Get(&comment, query)
	if err != nil {
		log.Printf("Error when fetching comment after create %s", err)
		return nil, errs.NewNotFoundError("Comment not found after being added")
	}

	return &comment, nil
}

func (db CommentRepositoryDb) DeleteComment(uuid string) *errs.AppError {
	var comment Comment

	query := fmt.Sprintf(`DELETE FROM comment WHERE uuid = '%s'`, uuid)

	_, err := db.client.NamedExec(query, comment)
	if err != nil {
		log.Println("Error when deleting comment", err)
		return errs.NewUnexpectedError("Something went wrong when deleting your comment")
	}

	return nil
}
