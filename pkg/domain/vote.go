package domain

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"

	"github.com/google/uuid"
)

type (
	Vote struct {
		UserId      uuid.UUID `db:"userId"`
		CountrySlug string    `db:"countrySlug"`
		Costume     uint8     `db:"costume"`
		Song        uint8     `db:"song"`
		Performance uint8     `db:"performance"`
		Props       uint8     `db:"props"`
	}

	Result struct {
		CountrySlug string `db:"countrySlug"`
		Costume     int    `db:"costume_total"`
		Song        int    `db:"song_total"`
		Performance int    `db:"performance_total"`
		Props       int    `db:"props_total"`
		Total       int    `db:"total"`
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

//go:generate mockgen -source=vote.go -destination=../../mocks/domain/mockVoteRepository.go -package=domain eurovision/pkg/domain
type VoteRepository interface {
	CreateVote(dto.Vote) (*Vote, *errs.AppError)
	UpdateVote(dto.VoteSingle) (*Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*Vote, *errs.AppError)
	GetResults() (*[]Result, *errs.AppError)
	GetResultsByUser(userId string) (*[]Result, *errs.AppError)
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
