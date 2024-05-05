package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID             uuid.UUID  `db:"uuid"`
	UserId           uuid.UUID  `db:"userId"`
	Text             []byte     `db:"text"`
	FileName         string     `db:"fileName"`
	CreatedAt        time.Time  `db:"createdAt"`
	ReplyToID        uuid.UUID  `db:"replyTo_uuid"`
	ReplyToUserId    uuid.UUID  `db:"replyTo_userId"`
	ReplyToFileName  []byte     `db:"replyTo_fileName"`
	ReplyToText      []byte     `db:"replyTo_text"`
	ReplyToCreatedAt *time.Time `db:"replyTo_createdAt"`
}

type CommentRepository interface {
	FindAllComments() ([]Comment, *errs.AppError)
	CreateComment(dto.Comment) (*Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
	FindCommentsAfter(string) ([]Comment, *errs.AppError)
}

func (comment Comment) ToDto() dto.Comment {
	var replyTo *dto.Comment
	if comment.ReplyToID != uuid.Nil {
		replyTo = &dto.Comment{
			UUID:      comment.ReplyToID,
			UserId:    comment.ReplyToUserId,
			Text:      string(comment.ReplyToText),
			FileName:  string(comment.ReplyToFileName),
			CreatedAt: *comment.ReplyToCreatedAt,
			ReplyTo:   nil,
		}
	}
	return dto.Comment{
		UUID:      comment.UUID,
		UserId:    comment.UserId,
		Text:      string(comment.Text),
		FileName:  comment.FileName,
		ReplyTo:   replyTo,
		CreatedAt: comment.CreatedAt,
	}
}
