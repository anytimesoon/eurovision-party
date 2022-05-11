package domain

import (
	"github.com/google/uuid"
)

type Country struct {
	Id            uuid.UUID `json:"id"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Votes         []Vote    `json:"votes"`
	Comments      []Comment `json:"comments"`
	Participating bool      `json:"participating"`
}
