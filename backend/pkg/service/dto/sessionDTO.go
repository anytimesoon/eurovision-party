package dto

import (
	"github.com/google/uuid"
)

type Session struct {
	SessionToken string    `json:"token"`
	User         User      `json:"user"`
	Bot          uuid.UUID `json:"botId"`
}
