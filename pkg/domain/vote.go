package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"

	"github.com/google/uuid"
)

// import "github.com/google/uuid"

type Vote struct {
	UUID        uuid.UUID `db:"uuid"`
	UserId      uuid.UUID `db:"userId"`
	CountryId   uuid.UUID `db:"countryId"`
	Costume     uint8     `db:"costume"`
	Song        uint8     `db:"song"`
	Performance uint8     `db:"performance"`
	Props       uint8     `db:"props"`
}

type VoteRepository interface {
	CreateVote(*dto.Vote) (*Vote, *errs.AppError)
	UpdateVote(*dto.Vote) (*Vote, *errs.AppError)
}

func (vote Vote) ToDto() dto.Vote {
	return dto.Vote{
		UUID:        vote.UUID,
		UserId:      vote.UserId,
		CountryId:   vote.CountryId,
		Costume:     vote.Costume,
		Song:        vote.Song,
		Performance: vote.Performance,
		Props:       vote.Props,
	}
}
