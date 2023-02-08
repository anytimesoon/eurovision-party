package dto

import (
	"eurovision/pkg/enum"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	Token      string
	Expiration time.Time
	UserId     uuid.UUID
	AuthLvl    enum.AuthLvl
}

type NewUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Slug  string    `json:"slug"`
	UUID  uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type EAuth struct {
	EToken string `json:"token"`
}

type AuthAndToken struct {
	Token   string
	AuthLvl enum.AuthLvl
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
