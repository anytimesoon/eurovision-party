package domain

import (
	"eurovision/pkg/dto"

	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID `db:"uuid"`
	AuthLvl AuthLvl   `db:"authLvl"`
	Name    string    `db:"name"`
	Slug    string    `db:"slug"`
	Icon    string    `db:"icon"`
}

type AuthLvl int

const (
	None AuthLvl = iota
	Admin
)

type UserRepository interface {
	FindAllUsers() ([]User, error)
	FindOneUser(string) (User, error)
	CreateUser(dto.User) (User, error)
	UpdateUser(dto.User) (User, error)
	DeleteUser(string) error
}

func (user User) ToDto() dto.User {
	return dto.User{
		Name: user.Name,
		Slug: user.Slug,
		UUID: user.UUID,
	}
}
