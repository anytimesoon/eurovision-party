package domain

import "github.com/google/uuid"

type Comment struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	CountryId uuid.UUID `json:"countryId"`
	Text      string
}
