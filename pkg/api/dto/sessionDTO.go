package dto

import (
	"github.com/google/uuid"
)

type CookieOpts struct {
	Path     string `json:"path"`
	MaxAge   int    `json:"maxAge"`
	Secure   bool   `json:"secure"`
	HttpOnly bool
	SameSite string
	Domain   string
}

// SessionAuth gets returned to users when they log in
type SessionAuth struct {
	Name         string
	SessionToken string    `json:"token"`
	User         User      `json:"user"`
	Bot          uuid.UUID `json:"botId"`
}
