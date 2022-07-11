package domain

import (
	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID `json:"id"`
	AuthLvl AuthLvl   `json:"authLvl"`
	Name    string    `json:"name"`
	Slug    string    `json:"slug"`
	Icon    string    `json:"icon"`
	// Votes    []Vote    `json:"votes"`
	// Comments []Comment `json:"comments"`
}

type AuthLvl int

const (
	None AuthLvl = iota
	Admin
)
