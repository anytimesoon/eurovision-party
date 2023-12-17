package dto

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"log"
	"net/http"
	"regexp"
	"strings"
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
		CookieOpts: CookieOpts{
			Path:     "/",
			MaxAge:   60 * 60 * 24 * 7,
			Secure:   false,
			HttpOnly: false,
			SameSite: 3,
		},
		User: user,
		Bot:  conf.App.BotId,
	}
}

type NewUser struct {
	Name  string    `json:"name"`
	Slug  string    `json:"slug"`
	UUID  uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type CookieOpts struct {
	Path     string `json:"path"`
	MaxAge   int    `json:"maxAge"`
	Secure   bool   `json:"secure"`
	HttpOnly bool
	SameSite http.SameSite
}

// SessionAuth gets returned to users when they log in
type SessionAuth struct {
	Name         string
	SessionToken string     `json:"token"`
	CookieOpts   CookieOpts `json:"opts"`
	User         User       `json:"user"`
	Bot          uuid.UUID  `json:"botId"`
}

func (nu *NewUser) Slugify() {
	re, err := regexp.Compile(`[[:^alnum:]]`)
	if err != nil {
		log.Fatal(err)
	}

	splitName := strings.Split(nu.Name, " ")
	finalName := make([]string, 0)
	for _, word := range splitName {
		word = re.ReplaceAllString(word, "")
		if word != "" {
			finalName = append(finalName, strings.ToLower(word))
		}
	}

	nu.Slug = strings.Join(finalName, "-")
}
