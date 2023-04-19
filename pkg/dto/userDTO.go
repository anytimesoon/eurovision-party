package dto

import (
	"eurovision/pkg/enum"
	"eurovision/pkg/errs"
	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID    `json:"id"`
	Name    string       `json:"name"`
	Slug    string       `json:"slug"`
	Icon    string       `json:"icon"`
	AuthLvl enum.AuthLvl `json:"authLvl"`
}

type UserImage struct {
	UUID  uuid.UUID `json:"id"`
	Image string    `json:"img"`
	//ImageBin image.Image
	AuthLvl     enum.AuthLvl `json:"authLvl"`
	TopLeft     int
	TopRight    int
	BottomLeft  int
	BottomRight int
}

func (u User) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(u.Name, "Name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
