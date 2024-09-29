package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/google/uuid"
)

type (
	Vote struct {
		UserId      uuid.UUID `boltholdIndex:"UserId"`
		CountrySlug string    //`boltholdKey:"CountrySlug"`
		Costume     uint8
		Song        uint8
		Performance uint8
		Props       uint8
		Total       int
	}

	Result struct {
		CountrySlug string
		Costume     int
		Song        int
		Performance int
		Props       int
		Total       int
	}
)

func (r Result) ToDto() dto.Result {
	return dto.Result{
		CountrySlug: r.CountrySlug,
		Costume:     r.Costume,
		Song:        r.Song,
		Performance: r.Performance,
		Props:       r.Props,
		Total:       r.Total,
	}
}

func (vote Vote) ToDto() dto.Vote {
	return dto.Vote{
		UserId:      vote.UserId,
		CountrySlug: vote.CountrySlug,
		Costume:     vote.Costume,
		Song:        vote.Song,
		Performance: vote.Performance,
		Props:       vote.Props,
	}
}

func (vote Vote) FromDTO(voteDTO dto.Vote) Vote {
	return Vote{
		UserId:      voteDTO.UserId,
		CountrySlug: voteDTO.CountrySlug,
		Costume:     voteDTO.Costume,
		Song:        voteDTO.Song,
		Performance: voteDTO.Performance,
		Props:       voteDTO.Props,
	}
}

func (vote Vote) ToResult() Result {
	return Result{
		CountrySlug: vote.CountrySlug,
		Costume:     int(vote.Costume),
		Song:        int(vote.Song),
		Performance: int(vote.Performance),
		Props:       int(vote.Props),
		Total:       vote.Total,
	}
}
