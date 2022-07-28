package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID `db:"uuid"`
	UserId    uuid.UUID `db:"userId"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"createdAt"`
}

type CommentRepository interface {
	FindAllComments() ([]Comment, *errs.AppError)
	CreateComment(dto.Comment) (*Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
}

func (comment Comment) ToDto() dto.Comment {
	return dto.Comment{
		UUID:      comment.UUID,
		UserId:    comment.UserId,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}
