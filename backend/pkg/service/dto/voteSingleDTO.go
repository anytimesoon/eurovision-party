package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
)

type VoteSingle struct {
	UserId      uuid.UUID       `json:"userId"`
	CountrySlug string          `json:"countrySlug"`
	Cat         enum.Categories `json:"cat"`
	Score       uint8           `json:"score"`
}

func (v VoteSingle) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(v.UserId.String(), "User ID")
	if message != "" {
		messages = append(messages, "You're not a user? We're as confused as you")
	}

	message = isPresent(v.CountrySlug, "Country ID")
	if message != "" {
		messages = append(messages, "Can't find the country you're trying to vote on")
	}

	message = isWithinRange(v.Score, v.Cat.String())
	if message != "" {
		messages = append(messages, message)
	}

	message = isValidCat(v.Cat.String())
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
