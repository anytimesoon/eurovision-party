package domain

import (
	"crypto/rand"
	"encoding/hex"
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	Token      string       `db:"token"`
	UserId     uuid.UUID    `db:"userId"`
	Expiration time.Time    `db:"texp"`
	EToken     string       `db:"etoken"`
	ETokExp    time.Time    `db:"etexp"`
	AuthLvl    enum.AuthLvl `db:"authLvl"`
	Slug       string       `db:"slug"`
}

func (a *Auth) GenerateSecureToken(len int) {
	a.Token = generateToken(len)
}

func (a *Auth) GenerateSecureEToken(len int) {
	a.EToken = generateToken(len)
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

func (a Auth) ToDTO() dto.Auth {
	return dto.Auth{
		Token:      a.EToken,
		Expiration: a.ETokExp,
		UserId:     a.UserId,
		AuthLvl:    a.AuthLvl,
	}
}
