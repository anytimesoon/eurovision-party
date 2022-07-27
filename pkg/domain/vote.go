package domain

import (
	"eurovision/pkg/dto"

	"github.com/google/uuid"
)

// import "github.com/google/uuid"

type Vote struct {
	UUID        uuid.UUID `db:"id"`
	UserId      uuid.UUID `db:"userId"`
	CountryId   uuid.UUID `db:"countryId"`
	Costume     uint8     `db:"costume"`
	Song        uint8     `db:"song"`
	Performance uint8     `db:"performance"`
	Props       uint8     `db:"props"`
}

type VoteRepository interface {
	CreateVote(dto.Vote) (Vote, error)
	// UpdateVote([]byte) (Vote, error)
}

func (vote Vote) ToDto() dto.Vote {
	return dto.Vote{}
}
