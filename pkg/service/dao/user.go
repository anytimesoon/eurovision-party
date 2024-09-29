package dao

import (
	dto2 "github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		UUID    uuid.UUID
		AuthLvl enum.AuthLvl `boltholdIndex:"AuthLvl"`
		Name    string
		Slug    string `boltholdUnique:"UniqueSlug"`
		Icon    string
	}
	NewUser struct {
		UUID    uuid.UUID
		AuthLvl enum.AuthLvl
		Name    string
		Slug    string
		Token   string
	}
)

func (user *User) ToDto() dto2.User {
	return dto2.User{
		Name:    user.Name,
		Slug:    user.Slug,
		UUID:    user.UUID,
		Icon:    user.Icon,
		AuthLvl: user.AuthLvl,
	}
}

func (user *User) FromDTO(userDTO dto2.User) *User {
	return &User{
		UUID:    userDTO.UUID,
		AuthLvl: userDTO.AuthLvl,
		Name:    userDTO.Name,
		Slug:    userDTO.Slug,
		Icon:    userDTO.Icon,
	}
}

func (user *User) ToNewUser() *NewUser {
	return &NewUser{
		UUID:    user.UUID,
		AuthLvl: user.AuthLvl,
		Name:    user.Name,
		Slug:    user.Slug,
		Token:   "",
	}
}

func (nu *NewUser) GenerateAuth() Auth {
	a := Auth{
		AuthToken:    nu.Token,
		UserId:       nu.UUID,
		AuthTokenExp: time.Now().Add(7 * 24 * time.Hour),
		AuthLvl:      nu.AuthLvl,
		Slug:         nu.Slug,
	}
	a.GenerateSecureToken(40)
	nu.Token = a.AuthToken
	return a
}
func (nu *NewUser) ToDTO() *dto2.NewUser {
	return &dto2.NewUser{
		Name: nu.Name,
		Slug: nu.Slug,
		UUID: nu.UUID,

		Token: nu.Token,
	}
}

func (nu *NewUser) FromDTO(nuDTO dto2.NewUser) {
	nu.Name = nuDTO.Name
	nu.Slug = nuDTO.Slug
	nu.UUID = uuid.New()
	nu.Token = nuDTO.Token
}

func (nu *NewUser) ToUser() *User {
	return &User{
		UUID:    nu.UUID,
		Name:    nu.Name,
		Slug:    nu.Slug,
		Icon:    "default",
		AuthLvl: nu.AuthLvl,
	}
}
