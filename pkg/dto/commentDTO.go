package dto

import (
	"eurovision/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c Comment) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(c.UserId.String(), "User ID")
	if message != "" {
		messages = append(messages, "You're not a user? We're as confused as you")
	}

	message = isPresent(c.Text, "Message body")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
