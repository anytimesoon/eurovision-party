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
	AuthToken       string `boltholdKey:"AuthToken"`
	UserId          uuid.UUID
	AuthTokenExp    time.Time
	SessionToken    string
	SessionTokenExp time.Time
	AuthLvl         enum.AuthLvl
	LastUpdated     time.Time
	Slug            string
}

type Session struct {
	AuthToken    string
	SessionToken string `boltholdKey:"SessionToken"`
	UserId       uuid.UUID
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
		Token:      a.SessionToken,
		Expiration: a.SessionTokenExp,
		UserId:     a.UserId,
		AuthLvl:    a.AuthLvl,
	}
}

// ToModifiedDTO takes in a new token, which is an encrypted version of an auth struct.
// A modified Auth is what we will send to the client
func (a *Auth) ToModifiedDTO(token string) dto.Auth {
	return dto.Auth{
		Token:      token,
		Expiration: a.SessionTokenExp,
		UserId:     a.UserId,
		AuthLvl:    a.AuthLvl,
	}
}
