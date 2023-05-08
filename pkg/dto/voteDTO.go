package dto

import (
	"github.com/google/uuid"
)

type Vote struct {
	UserId      uuid.UUID `json:"userId"`
	CountrySlug string    `json:"countrySlug"`
	Costume     uint8     `json:"costume"`
	Song        uint8     `json:"song"`
	Performance uint8     `json:"performance"`
	Props       uint8     `json:"props"`
}
