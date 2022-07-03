package domain

import (
	"github.com/google/uuid"
)

type Country struct {
	UUID          uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Votes         []Vote    `json:"votes"`
	Comments      []Comment `json:"comments"`
	Participating bool      `json:"participating"`
}
