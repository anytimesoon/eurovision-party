package dto

import (
	"log"
	"regexp"
	"strings"

	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/google/uuid"
)

type NewUser struct {
	Name      string          `json:"name"`
	Slug      string          `json:"slug"`
	UUID      uuid.UUID       `json:"id"`
	AuthLvl   authLvl.AuthLvl `json:"authLvl"`
	Token     string          `json:"token"`
	CreatedBy uuid.UUID       `json:"createdBy"`
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
