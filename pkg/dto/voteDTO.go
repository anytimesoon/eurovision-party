package dto

import "github.com/google/uuid"

type Vote struct {
	UUID        uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	CountryId   uuid.UUID `json:"countryId"`
	Costume     uint8     `json:"costume"`
	Song        uint8     `json:"song"`
	Performance uint8     `json:"performance"`
	Props       uint8     `json:"props"`
}
