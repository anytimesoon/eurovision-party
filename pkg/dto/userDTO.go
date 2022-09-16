package dto

import (
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
	Icon string    `json:"icon"`
}

func (u User) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(u.Name, "Name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
