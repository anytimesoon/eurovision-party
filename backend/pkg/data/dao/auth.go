package dao

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
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

func (a *Auth) ToDTO() dto2.Auth {
	return dto2.Auth{
		Token:      a.SessionToken,
		Expiration: a.SessionTokenExp,
		UserId:     a.UserId,
		AuthLvl:    a.AuthLvl,
	}
}

func (a *Auth) ToSession(token string, user *User) *dto2.Session {
	return &dto2.Session{
		SessionToken: token,
		User:         user.ToDto(),
		Bot:          conf.App.BotId,
	}
}
