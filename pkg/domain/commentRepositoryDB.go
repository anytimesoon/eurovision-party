package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

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

	newCommentId := uuid.New().String()

	createCommentQuery := "INSERT INTO comment(uuid, userId, text) VALUES (?, ?, ?)"
	getCommentQuery := "SELECT * FROM comment WHERE uuid = ?"

	// Begin transaction that will create a new comment then return the new comment
	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction for new comment for user %s, %s", commentDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "you comment")
	}

	_, err = tx.Exec(createCommentQuery, newCommentId, commentDTO.UserId, commentDTO.Text)
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", commentDTO.UserId, err)
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your comment")
	}

	err = tx.Get(&comment, getCommentQuery, newCommentId)
	if err != nil {
		log.Printf("Error when fetching comment after create %s", err)
		_ = tx.Rollback()
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your comment")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error when commiting comment transaction for user %s. %s", commentDTO.UUID, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "a new comment")
	}

	return &comment, nil
}

func (db CommentRepositoryDb) DeleteComment(uuid string) *errs.AppError {
	var comment Comment

	query := "DELETE FROM comment WHERE uuid = ?"

	_, err := db.client.Exec(query, uuid)
	if err != nil {
		log.Println("Error when deleting comment", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "your comment")
	}

	return nil
}
