package dto

import "github.com/google/uuid"

type Vote struct {
	UUID        uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	CountryId   uuid.UUID `json:"countryId"`
	Costume     int8      `json:"costume"`
	Song        int8      `json:"song"`
	Performance int8      `json:"performance"`
	Props       int8      `json:"props"`
}
