package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"log"

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

	query := `SELECT
					c1.uuid,
					c1.text,
					c1.userId,
					c1.createdAt,
					c2.uuid as replyTo_uuid,
					c2.userId as replyTo_userId,
					c2.text as replyTo_text,
					c2.createdAt as replyTo_createdAt
				FROM comment c1 
					LEFT JOIN comment c2 ON c1.replyTo = c2.uuid
				ORDER BY c1.createdAt`
	err := db.client.Select(&comments, query)
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	return comments, nil
}

func (db CommentRepositoryDb) FindCommentsAfter(commentId string) ([]Comment, *errs.AppError) {
	comments := make([]Comment, 0)

	query := `SELECT
					c1.uuid,
					c1.text,
					c1.userId,
					c1.createdAt,
					c2.uuid as replyTo_uuid,
					c2.userId as replyTo_userId,
					c2.text as replyTo_text,
					c2.createdAt as replyTo_createdAt
				FROM comment c1 
					LEFT JOIN comment c2 ON c1.replyTo = c2.uuid
				WHERE c1.createdAt > (SELECT createdAt
				                      FROM comment
				                      WHERE uuid = ?)
				ORDER BY c1.createdAt`
	err := db.client.Select(&comments, query, commentId)
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	return comments, nil
}

func (db CommentRepositoryDb) CreateComment(commentDTO dto.Comment) (*Comment, *errs.AppError) {
	var comment Comment

	createCommentQuery := "INSERT INTO comment(uuid, userId, text) VALUES (?, ?, ?)"
	createCommentWithReplyQuery := "INSERT INTO comment(uuid, userId, text, replyTo) VALUES (?, ?, ?, ?)"
	getCommentQuery := `SELECT
							c1.uuid,
							c1.text,
							c1.userId,
							c1.createdAt,
							c2.uuid as replyTo_uuid,
							c2.userId as replyTo_userId,
							c2.text as replyTo_text,
							c2.createdAt as replyTo_createdAt
						FROM comment c1 
							LEFT JOIN comment c2 ON c1.replyTo = c2.uuid
						WHERE c1.uuid = ?`

	// Begin transaction that will create a new comment then return the new comment
	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting transaction for new comment for user %s, %s", commentDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "you comment")
	}

	if commentDTO.ReplyTo != nil {
		_, err = tx.Exec(createCommentWithReplyQuery, commentDTO.UUID.String(), commentDTO.UserId, commentDTO.Text, commentDTO.ReplyTo.UUID)
	} else {
		_, err = tx.Exec(createCommentQuery, commentDTO.UUID.String(), commentDTO.UserId, commentDTO.Text)
	}
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", commentDTO.UserId, err)
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your comment")
	}

	err = tx.Get(&comment, getCommentQuery, commentDTO.UUID.String())
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
	query := "DELETE FROM comment WHERE uuid = ?"

	_, err := db.client.Exec(query, uuid)
	if err != nil {
		log.Println("Error when deleting comment", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "your comment")
	}

	return nil
}
