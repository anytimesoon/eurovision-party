package domain

import "github.com/google/uuid"

type Country struct {
	UUID          uuid.UUID `db:"uuid"`
	Name          string    `db:"name"`
	Slug          string    `db:"slug"`
	BandName      string    `db:"bandName"`
	SongName      string    `db:"songName"`
	Flag          string    `db:"flag"`
	Participating bool      `db:"participating"`
}

type CountryRepository interface {
	FindAllCountries() ([]Country, error)
}
