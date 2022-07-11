package dto

import "github.com/google/uuid"

type Comment struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	CountryId uuid.UUID `json:"countryId"`
	Text      string
}
