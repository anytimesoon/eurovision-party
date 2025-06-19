package dao

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID
	AuthLvl enum.AuthLvl `boltholdIndex:"AuthLvl"`
	Name    string
	Slug    string `boltholdUnique:"UniqueSlug"`
	Icon    string
}

func (u User) ToDto() dto.User {
	return dto.User{
		Name:    u.Name,
		Slug:    u.Slug,
		UUID:    u.UUID,
		Icon:    u.Icon,
		AuthLvl: u.AuthLvl,
	}
}

func (u User) FromDTO(userDTO dto.User) *User {
	return &User{
		UUID:    userDTO.UUID,
		AuthLvl: userDTO.AuthLvl,
		Name:    userDTO.Name,
		Slug:    userDTO.Slug,
		Icon:    userDTO.Icon,
	}
}

func (u User) ToNewUserDTO(auth Auth) *dto.NewUser {
	return &dto.NewUser{
		UUID:    u.UUID,
		AuthLvl: u.AuthLvl,
		Name:    u.Name,
		Slug:    u.Slug,
		Token:   auth.AuthToken,
	}
}

func (u User) FromNewUserDTO(newUser dto.NewUser) *User {
	return &User{
		UUID:    uuid.New(),
		AuthLvl: newUser.AuthLvl,
		Name:    newUser.Name,
		Slug:    newUser.Slug,
		Icon:    "default",
	}
}

func (u User) GenerateAuth() Auth {
	a := Auth{
		UserId:       u.UUID,
		AuthTokenExp: time.Now().Add(7 * 24 * time.Hour),
		AuthLvl:      u.AuthLvl,
		Slug:         u.Slug,
	}
	a.GenerateSecureToken(40)
	return a
}
