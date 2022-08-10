package dto

import (
	"eurovision/pkg/errs"
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
	Icon string    `json:"icon"`
}

func (u User) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(u.Name, "Name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}

func (u *User) Slugify() {
	re, err := regexp.Compile(`[[:^alnum:]]`)
	if err != nil {
		log.Fatal(err)
	}

	splitName := strings.Split(u.Name, " ")
	finalName := make([]string, 0)
	for _, word := range splitName {
		word = re.ReplaceAllString(word, "")
		if word != "" {
			finalName = append(finalName, strings.ToLower(word))
		}
	}

	u.Slug = strings.Join(finalName, "-")
}
