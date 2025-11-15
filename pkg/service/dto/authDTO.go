package dto

import (
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"

	"github.com/google/uuid"
)

// Auth is used internally to verify a user, auth level, and expiration of session.
// It should never be returned to a user
type Auth struct {
	Token      string
	Expiration time.Time
	UserId     uuid.UUID
	AuthLvl    authLvl.AuthLvl
}

func (a Auth) ToSession(user User) Session {
	return Session{
		SessionToken: a.Token,
		User:         user,
		Bot:          conf.App.BotId,
	}
}
