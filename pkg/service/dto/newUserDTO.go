package dto

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/google/uuid"
	"log"
	"regexp"
	"strings"
)

type NewUser struct {
	Name    string       `json:"name"`
	Slug    string       `json:"slug"`
	UUID    uuid.UUID    `json:"id"`
	AuthLvl enum.AuthLvl `json:"authLvl"`
	Token   string       `json:"token"`
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
