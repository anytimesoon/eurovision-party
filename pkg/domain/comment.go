package domain

import (
	"eurovision/pkg/dto"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentRepository interface {
	FindAllComments() ([]Comment, error)
	CreateComment(dto.Comment) (Comment, error)
	DeleteComment(string) error
}

func (comment Comment) ToDto() dto.Comment {
	return dto.Comment{
		UUID:      comment.UUID,
		UserId:    comment.UserId,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}
