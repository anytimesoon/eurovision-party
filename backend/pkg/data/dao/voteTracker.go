package dao

import (
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
)

type VoteTracker struct {
	Count           int
	HasBeenNotified bool
	Country         Country
}

func (v VoteTracker) ToDto() dto.VoteTracker {
	return dto.VoteTracker{
		Count:   v.Count,
		Country: v.Country.ToDto(),
	}
}
