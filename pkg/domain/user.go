package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"

	"github.com/google/uuid"
)

type (
	User struct {
		UUID    uuid.UUID    `db:"uuid"`
		AuthLvl enum.AuthLvl `db:"authLvl"`
		Name    string       `db:"name"`
		Slug    string       `db:"slug"`
		Icon    string       `db:"icon"`
	}
	NewUser struct {
		UUID  uuid.UUID `db:"uuid"`
		Name  string    `db:"name"`
		Slug  string    `db:"slug"`
		Token string    `db:"authToken"`
	}
)

//go:generate mockgen -source=user.go -destination=../../mocks/domain/mockUserRepository.go -package=domain eurovision/pkg/domain
type UserRepository interface {
	FindAllUsers() ([]User, *errs.AppError)
	FindOneUser(string) (*User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	FindRegisteredUsers() (*[]NewUser, *errs.AppError)
	UpdateUser(dto.User) (*User, *dto.Comment, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*User, *dto.Comment, *errs.AppError)
}

func (user User) ToDto() dto.User {
	return dto.User{
		Name:    user.Name,
		Slug:    user.Slug,
		UUID:    user.UUID,
		Icon:    user.Icon,
		AuthLvl: user.AuthLvl,
	}
}

func (user *NewUser) ToDTO() *dto.NewUser {
	return &dto.NewUser{
		Name:  user.Name,
		Slug:  user.Slug,
		UUID:  user.UUID,
		Token: user.Token,
	}
}
