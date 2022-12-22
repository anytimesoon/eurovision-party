package dto

import (
	"eurovision/pkg/errs"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Country struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Participating bool      `json:"participating"`
}

func (c Country) Validate() *errs.AppError {
	messages := make([]string, 0)

	message := isPresent(c.BandName, "Band name")
	if message != "" {
		messages = append(messages, message)
	}

	message = isPresent(c.SongName, "Song name")
	if message != "" {
		messages = append(messages, message)
	}

	return messagesToError(messages)
}

func (c Country) GetName() string {
	return c.Name
}

func (c *Country) SetSlug(slug string) {
	c.Slug = slug
}

func (c Country) VerifySlug(slug string, db *sqlx.DB) bool {
	query := fmt.Sprintf("SELECT * FROM country WHERE slug = '%s'", slug)
	row, err := db.Queryx(query)
	if err != nil {
		return false
	}

	if row.Rows != nil {
		return false
	}
	return true
}
