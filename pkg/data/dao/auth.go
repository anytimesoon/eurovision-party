package dao

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"time"
)

type Auth struct {
	AuthToken       string `boltholdKey:"AuthToken"`
	UserId          uuid.UUID
	AuthTokenExp    time.Time
	SessionToken    string
	SessionTokenExp time.Time
	AuthLvl         authLvl.AuthLvl
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

func (a *Auth) ToSession(token string, user *User) *dto.Session {
	return &dto.Session{
		SessionToken: token,
		User:         user.ToDto(),
		Bot:          conf.App.BotId,
	}
}
