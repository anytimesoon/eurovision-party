package domain

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	AuthToken    string       `db:"authToken"`
	UserId       uuid.UUID    `db:"userId"`
	AuthTokenExp time.Time    `db:"authTokenExp"`
	SessionToken string       `db:"sessionToken"`
	AuthLvl      enum.AuthLvl `db:"authLvl"`
	LastUpdated  time.Time    `db:"lastUpdated"`
	Slug         string       `db:"slug"`
}

func (a *Auth) GenerateSecureToken(len int) {
	a.AuthToken = generateToken(len)
}

func (a *Auth) GenerateSecureSessionToken(len int) {
	a.SessionToken = generateToken(len)
}

func generateToken(len int) string {
	b := make([]byte, len)
	isNotGenerated := true

	for isNotGenerated {
		_, err := rand.Read(b)
		if err != nil {
			err = nil
		} else {
			isNotGenerated = false
		}
	}

	return hex.EncodeToString(b)
}

func (a *Auth) ToDTO() dto.Auth {
	return dto.Auth{
		Token:   a.SessionToken,
		UserId:  a.UserId,
		AuthLvl: a.AuthLvl,
	}
}

// ToModifiedDTO takes in a new token, which is an encrypted version of an auth struct.
// A modified Auth is what we will send to the client
func (a *Auth) ToModifiedDTO(token string) dto.Auth {
	return dto.Auth{
		Token:   token,
		UserId:  a.UserId,
		AuthLvl: a.AuthLvl,
	}
}
