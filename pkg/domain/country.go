package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
)

type Country struct {
	Name          string
	Slug          string `boltholdKey:"Slug" boltholdUnique:"UniqueSlug"`
	BandName      string
	SongName      string
	Flag          string
	Participating bool
}

type CountryRepository interface {
	FindAllCountries() (*[]Country, *errs.AppError)
	FindOneCountry(string) (*Country, *errs.AppError)
	UpdateCountry(dto.Country) (*Country, *errs.AppError)
	FindParticipating() (*[]Country, *errs.AppError)
}

func (c Country) ToDto() dto.Country {
	return dto.Country{
		Name:          c.Name,
		Slug:          c.Slug,
		BandName:      c.BandName,
		SongName:      c.SongName,
		Flag:          c.Flag,
		Participating: c.Participating,
	}
}

func (c Country) FromDTO(dto dto.Country) Country {
	return Country{
		Name:          dto.Name,
		Slug:          dto.Slug,
		BandName:      dto.BandName,
		SongName:      dto.SongName,
		Flag:          dto.Flag,
		Participating: dto.Participating,
	}
}
