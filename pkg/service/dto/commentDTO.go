package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID               uuid.UUID `json:"id"`
	UserId             uuid.UUID `json:"userId"`
	Text               string    `json:"text"`
	FileName           string    `json:"fileName"`
	CreatedAt          time.Time `json:"createdAt"`
	ReplyTo            *Comment  `json:"replyToComment,omitempty"`
	IsVoteNotification bool      `json:"isVoteNotification"`
}

func (c Comment) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(c.UserId.String(), "User ID")
	if message != "" {
		messages = append(messages, "You're not a user? We're as confused as you")
	}

	commentTextMessage := isPresent(c.Text, "Comment text")
	commentImageMessage := isPresent(c.FileName, "Comment image")
	if commentTextMessage != "" && commentImageMessage != "" {
		messages = append(messages, commentTextMessage, commentImageMessage)
	}

	return messagesToError(messages)
}
