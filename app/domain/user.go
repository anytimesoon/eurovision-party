package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	AuthLvl  AuthLvl   `json:"authLvl"`
	Name     string    `json:"Name"`
	Icon     string    `json:"icon"`
	Votes    []Vote    `json:"votes"`
	Comments []Comment `json:"comments"`
}

type AuthLvl int

const (
	None AuthLvl = iota
	Admin
)
