package dto

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"net/http"
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

func (a Auth) ToSession(user User, cookie http.Cookie) SessionAuth {
	return SessionAuth{
		Name:         "session",
		SessionToken: a.Token,
		CookieOpts: CookieOpts{
			Path:     cookie.Path,
			MaxAge:   cookie.MaxAge,
			Secure:   cookie.Secure,
			HttpOnly: cookie.HttpOnly,
			SameSite: sameSiteToString(cookie.SameSite),
			Domain:   cookie.Domain,
		},
		User: user,
		Bot:  conf.App.BotId,
	}
}
