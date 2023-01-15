package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID    `db:"uuid"`
	AuthLvl enum.AuthLvl `db:"authLvl"`
	Name    string       `db:"name"`
	Email   string       `db:"email"`
	Slug    string       `db:"slug"`
	Icon    string       `db:"icon"`
}

//go:generate mockgen -source=user.go -destination=../../mocks/domain/mockUserRepository.go -package=domain eurovision/pkg/domain
type UserRepository interface {
	FindAllUsers() ([]User, *errs.AppError)
	FindOneUser(string) (*User, *errs.AppError)
	UpdateUser(dto.User) (*User, *errs.AppError)
	DeleteUser(string) *errs.AppError
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
