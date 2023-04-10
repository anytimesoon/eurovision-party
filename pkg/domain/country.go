package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"

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

//go:generate mockgen -source=country.go -destination=../../mocks/domain/mockCountryRepository.go -package=domain eurovision/pkg/domain
type CountryRepository interface {
	FindAllCountries() (*[]Country, *errs.AppError)
	FindOneCountry(string) (*Country, *errs.AppError)
	UpdateCountry(dto.Country) (*Country, *errs.AppError)
	FindParticipating() (*[]Country, *errs.AppError)
}

func (c Country) ToDto() dto.Country {
	return dto.Country{
		ID:            c.UUID,
		Name:          c.Name,
		Slug:          c.Slug,
		BandName:      c.BandName,
		SongName:      c.SongName,
		Flag:          c.Flag,
		Participating: c.Participating,
	}
}
