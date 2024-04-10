package dto

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/google/uuid"
	"log"
	"regexp"
	"strings"
)

// Auth is used internally to verify a user, auth level, and expiration of session.
// The token refers to a session token, not auth token.
// It should never be returned to a user
type Auth struct {
	Token   string //session token
	UserId  uuid.UUID
	AuthLvl enum.AuthLvl
}

func (a Auth) ToSession(user User) *SessionAuth {
	return &SessionAuth{
		SessionToken: a.Token,
		User:         user,
		Bot:          conf.App.BotId,
	}
}

type NewUser struct {
	Name  string    `json:"name"`
	Slug  string    `json:"slug"`
	UUID  uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

// SessionAuth gets returned to users when they log in
type SessionAuth struct {
	SessionToken string    `json:"token"`
	User         User      `json:"user"`
	Bot          uuid.UUID `json:"botId"`
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
