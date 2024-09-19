package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
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

type UserRepository interface {
	CreateUser(dto.NewUser) (*NewUser, *errs.AppError)
	FindAllUsers() ([]User, *errs.AppError)
	FindOneUser(string) (*User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	FindRegisteredUsers() (*[]NewUser, *errs.AppError)
	UpdateUser(dto.User) (*User, *dto.Comment, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*User, *dto.Comment, *errs.AppError)
	VerifySlug(*dto.NewUser) error
}

func (user *User) ToDto() dto.User {
	return dto.User{
		Name:    user.Name,
		Slug:    user.Slug,
		UUID:    user.UUID,
		Icon:    user.Icon,
		AuthLvl: user.AuthLvl,
	}
}

func (user *User) FromDTO(userDTO dto.User) *User {
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
func (nu *NewUser) ToDTO() *dto.NewUser {
	return &dto.NewUser{
		Name: nu.Name,
		Slug: nu.Slug,
		UUID: nu.UUID,

		Token: nu.Token,
	}
}

func (nu *NewUser) FromDTO(nuDTO dto.NewUser) {
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
