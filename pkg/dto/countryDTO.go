package dto

import (
	"github.com/google/uuid"
)

type CountryData struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Participating bool      `json:"participating"`
}

type Countries struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Country `json:"data"`
}

type Country struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    CountryData `json:"data"`
}
