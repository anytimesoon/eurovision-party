package domain

import (
	"eurovision/pkg/dto"

	"github.com/google/uuid"
)

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
	FindOneCountry(string) (Country, error)
	UpdateCountry(dto.CountryData) (Country, error)
}

func (c Country) ToDto() dto.Country {
	return dto.Country{
		Success: true,
		Data: dto.CountryData{
			ID:            c.UUID,
			Name:          c.Name,
			Slug:          c.Slug,
			BandName:      c.BandName,
			SongName:      c.SongName,
			Flag:          c.Flag,
			Participating: c.Participating,
		},
	}
}
