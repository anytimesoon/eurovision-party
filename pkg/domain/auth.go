package domain

import (
	"crypto/rand"
	"encoding/hex"
	"eurovision/pkg/dto"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	Token      string    `db:"token"`
	UserId     uuid.UUID `db:"userId"`
	Expiration time.Time `db:"expiration"`
	AuthLvl    AuthLvl   `db:"authLvl"`
	Slug       string    `db:"slug"`
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
		Token:      a.Token,
		UserId:     a.UserId,
		Expiration: a.Expiration,
	}
}
