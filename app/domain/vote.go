package domain

import "github.com/google/uuid"

type Vote struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	CountryId uuid.UUID `json:"countryId"`
	Value     int
}
