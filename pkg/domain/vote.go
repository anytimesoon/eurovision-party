package domain

import "github.com/google/uuid"

type Vote struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	CountryId   uuid.UUID `json:"countryId"`
	Costume     int       `json:"costume"`
	Song        int       `json:"song"`
	Performance int       `json:"performance"`
	Props       int       `json:"props"`
}
