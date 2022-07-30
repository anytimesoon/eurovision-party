package dto

import (
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type Country struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Participating bool      `json:"participating"`
}

func (c Country) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(c.BandName, "Band name")
	if message != "" {
		messages = append(messages, message)
	}

	message = isPresent(c.SongName, "Song name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
