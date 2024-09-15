package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID
	UserId    uuid.UUID
	Text      string
	FileName  string
	CreatedAt time.Time
	ReplyTo   *Comment
}

type CommentRepository interface {
	FindAllComments() ([]Comment, *errs.AppError)
	CreateComment(dto.Comment) (*Comment, *errs.AppError)
	DeleteComment(string) *errs.AppError
	FindCommentsAfter(string) ([]Comment, *errs.AppError)
}

func (comment *Comment) ToDto() dto.Comment {
	var replyTo *dto.Comment
	if comment.ReplyTo != nil {
		rtc := comment.ReplyTo
		replyTo = &dto.Comment{
			UUID:      rtc.UUID,
			UserId:    rtc.UserId,
			Text:      rtc.Text,
			FileName:  rtc.FileName,
			CreatedAt: rtc.CreatedAt,
			ReplyTo:   nil,
		}
	}
	return dto.Comment{
		UUID:      comment.UUID,
		UserId:    comment.UserId,
		Text:      comment.Text,
		FileName:  comment.FileName,
		ReplyTo:   replyTo,
		CreatedAt: comment.CreatedAt,
	}
}

func (comment *Comment) FromDTO(dto dto.Comment) {
	var replyTo *Comment
	if dto.ReplyTo != nil {
		rtc := dto.ReplyTo
		replyTo = &Comment{
			UUID:      rtc.UUID,
			UserId:    rtc.UserId,
			Text:      rtc.Text,
			FileName:  rtc.FileName,
			CreatedAt: rtc.CreatedAt,
			ReplyTo:   nil,
		}
	}

	comment.UUID = dto.UUID
	comment.UserId = dto.UserId
	comment.Text = dto.Text
	comment.FileName = dto.FileName
	comment.ReplyTo = replyTo
	comment.CreatedAt = time.Now()
}
