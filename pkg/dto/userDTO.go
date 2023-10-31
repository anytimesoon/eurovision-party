package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"image"
	"mime/multipart"
)

type User struct {
	UUID    uuid.UUID    `json:"id"`
	Name    string       `json:"name"`
	Slug    string       `json:"slug"`
	Icon    string       `json:"icon"`
	AuthLvl enum.AuthLvl `json:"authLvl"`
}

type UserAvatar struct {
	UUID    uuid.UUID
	File    multipart.File
	Header  *multipart.FileHeader
	CropBox image.Rectangle
}

func (u User) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(u.Name, "Name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}
