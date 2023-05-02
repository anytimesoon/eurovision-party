package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

type Vote struct {
	UUID        uuid.UUID `db:"uuid"`
	UserId      uuid.UUID `db:"userId"`
	CountrySlug string    `db:"countrySlug"`
	Costume     uint8     `db:"costume"`
	Song        uint8     `db:"song"`
	Performance uint8     `db:"performance"`
	Props       uint8     `db:"props"`
}

//go:generate mockgen -source=vote.go -destination=../../mocks/domain/mockVoteRepository.go -package=domain eurovision/pkg/domain
type VoteRepository interface {
	CreateVote(dto.Vote) (*Vote, *errs.AppError)
	UpdateVote(dto.VoteSingle) (*Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*Vote, *errs.AppError)
}

func (vote Vote) ToDto() dto.Vote {
	return dto.Vote{
		UUID:        vote.UUID,
		UserId:      vote.UserId,
		CountrySlug: vote.CountrySlug,
		Costume:     vote.Costume,
		Song:        vote.Song,
		Performance: vote.Performance,
		Props:       vote.Props,
	}
}
