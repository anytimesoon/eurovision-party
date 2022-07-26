package dto

import "github.com/google/uuid"

type User struct {
	UUID uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
	Icon string    `json:"icon"`
}
