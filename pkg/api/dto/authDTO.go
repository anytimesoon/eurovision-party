package dto

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"time"

	"github.com/google/uuid"
)

// Auth is used internally to verify a user, auth level, and expiration of session.
// The token refers to a session token, not auth token.
// It should never be returned to a user
type Auth struct {
	Token      string //session token
	Expiration time.Time
	UserId     uuid.UUID
	AuthLvl    enum.AuthLvl
}

func (a Auth) ToSession(user User) SessionAuth {
	return SessionAuth{
		Name:         "session",
		SessionToken: a.Token,
		User:         user,
		Bot:          conf.App.BotId,
	}
}
