package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID `db:"uuid"`
	AuthLvl AuthLvl   `db:"authLvl"`
	Name    string    `db:"name"`
	Slug    string    `db:"slug"`
	Icon    string    `db:"icon"`
}

type AuthLvl uint8

const (
	None AuthLvl = iota
	Admin
	Bot
)

//go:generate mockgen -source=user.go -destination=../../mocks/domain/mockUserRepository.go -package=domain eurovision/pkg/domain
type UserRepository interface {
	FindAllUsers() ([]User, *errs.AppError)
	FindOneUser(string) (*User, *errs.AppError)
	CreateUser(dto.User) (*User, *errs.AppError)
	UpdateUser(dto.User) (*User, *errs.AppError)
	DeleteUser(string) *errs.AppError
}

func (user User) ToDto() dto.User {
	return dto.User{
		Name: user.Name,
		Slug: user.Slug,
		UUID: user.UUID,
		Icon: user.Icon,
	}
}
