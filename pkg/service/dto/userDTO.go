package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"image"
	"mime/multipart"
)

type User struct {
	UUID      uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Slug      string          `json:"slug"`
	Icon      string          `json:"icon"`
	AuthLvl   authLvl.AuthLvl `json:"authLvl"`
	Invites   []uuid.UUID     `json:"invites"`
	CreatedBy uuid.UUID       `json:"createdBy"`
	CanInvite bool            `json:"canInvite"`
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
