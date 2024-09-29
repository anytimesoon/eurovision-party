package dto

import (
	"github.com/google/uuid"
	"net/http"
)

func sameSiteToString(sameSite http.SameSite) string {
	switch sameSite {
	case 2:
		return "lax"
	case 3:
		return "strict"
	default:
		return "none"
	}
}

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
	SessionToken string     `json:"token"`
	CookieOpts   CookieOpts `json:"opts"`
	User         User       `json:"user"`
	Bot          uuid.UUID  `json:"botId"`
}
