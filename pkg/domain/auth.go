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

func (a *Auth) GenerateSecureToken() {
	b := make([]byte, 80)
	isNotGenerated := true

	for isNotGenerated {
		_, err := rand.Read(b)
		if err != nil {
			err = nil
		} else {
			isNotGenerated = false
		}
	}

	a.Token = hex.EncodeToString(b)
}

func (a Auth) ToDTO() dto.Auth {
	return dto.Auth{
		Token:      a.EToken,
		UserId:     a.UserId,
		Expiration: a.ETokExp,
		AuthLvl:    a.AuthLvl,
	}
}
