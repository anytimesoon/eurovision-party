package dto

import (
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type Vote struct {
	UUID        uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	CountryId   uuid.UUID `json:"countryId"`
	Costume     uint8     `json:"costume"`
	Song        uint8     `json:"song"`
	Performance uint8     `json:"performance"`
	Props       uint8     `json:"props"`
}

func (v Vote) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(v.UserId.String(), "User ID")
	if message != "" {
		messages = append(messages, "You're not a user? We're as confused as you")
	}

	message = isPresent(v.CountryId.String(), "Country ID")
	if message != "" {
		messages = append(messages, "Can't find the country you're trying to vote on")
	}

	message = isWithinRange(v.Costume, "costume")
	if message != "" {
		messages = append(messages, message)
	}

	message = isWithinRange(v.Song, "song")
	if message != "" {
		messages = append(messages, message)
	}

	message = isWithinRange(v.Performance, "performance")
	if message != "" {
		messages = append(messages, message)
	}

	message = isWithinRange(v.Costume, "props")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
