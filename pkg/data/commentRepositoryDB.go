package data

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/timshannon/bolthold"
	"log"
	"sort"
)

type CommentRepository interface {
	FindAllComments() ([]dao.Comment, *errs.AppError)
	CreateComment(dto.Comment) (*dao.Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
	FindCommentsAfter(string) ([]dao.Comment, *errs.AppError)
}

type CommentRepositoryDb struct {
	store *bolthold.Store
}

func NewCommentRepositoryDb(store *bolthold.Store) CommentRepositoryDb {
	return CommentRepositoryDb{store}
}

func (db CommentRepositoryDb) FindAllComments() ([]dao.Comment, *errs.AppError) {
	comments := make([]dao.Comment, 0)

	err := db.store.Find(&comments, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	sort.SliceStable(comments, func(i, j int) bool {
		return comments[i].CreatedAt.Before(comments[j].CreatedAt)
	})

	return comments, nil
}

func (db CommentRepositoryDb) FindCommentsAfter(commentId string) ([]dao.Comment, *errs.AppError) {
	comments := make([]dao.Comment, 0)
	var latestComment dao.Comment

	err := db.store.Get(commentId, &latestComment)
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	err = db.store.Find(&comments, bolthold.Where("CreatedAt").Gt(latestComment.CreatedAt))
	if err != nil {
		log.Println("Error while querying comment table", err)
		return nil, errs.NewNotFoundError("No comments found")
	}

	sort.SliceStable(comments, func(i, j int) bool {
		return comments[i].CreatedAt.Before(comments[j].CreatedAt)
	})

	return comments, nil
}

func (db CommentRepositoryDb) CreateComment(commentDTO dto.Comment) (*dao.Comment, *errs.AppError) {
	var comment dao.Comment

	comment.FromDTO(commentDTO)
	err := db.store.Insert(comment.UUID.String(), &comment)
	if err != nil {
		log.Printf("Error when creating comment from user %s, %s", commentDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your comment")
	}

	return &comment, nil
}

func (db CommentRepositoryDb) DeleteComment(uuid string) *errs.AppError {
	var comment dao.Comment

	err := db.store.Delete(uuid, comment)
	if err != nil {
		log.Println("Error when deleting comment", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "your comment")
	}

	return nil
}
