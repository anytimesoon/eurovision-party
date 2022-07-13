package dto

import "github.com/google/uuid"

type VoteData struct {
	UUID        uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	CountryId   uuid.UUID `json:"countryId"`
	Costume     int       `json:"costume"`
	Song        int       `json:"song"`
	Performance int       `json:"performance"`
	Props       int       `json:"props"`
}

type Vote struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    VoteData `json:"data"`
}

type Votes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []VoteData `json:"data"`
}
