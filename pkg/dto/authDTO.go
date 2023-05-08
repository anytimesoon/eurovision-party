package dto

import (
	"eurovision/pkg/enum"
	"log"
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

func (a Auth) ToSession() SessionAuth {
	return SessionAuth{
		SessionToken: a.Token,
		Exp:          a.Expiration,
	}
}

type NewUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Slug  string    `json:"slug"`
	UUID  uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

// SessionAuth gets returned to users if their session is valid
type SessionAuth struct {
	SessionToken string    `json:"token"`
	Exp          time.Time `json:"exp"`
}

// AuthAndToken is used interally to verify the authorization level of a user.
// It should never be sent to a user. It appears to be the same as Auth.
// TODO: can probably be salely replaced with Auth
//type AuthAndToken struct {
//	Token   string
//	AuthLvl enum.AuthLvl
//	UUID    uuid.UUID
//}

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
