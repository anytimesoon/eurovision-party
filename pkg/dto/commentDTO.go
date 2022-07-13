package dto

import (
	"time"

	"github.com/google/uuid"
)

type CommentData struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Comments struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []CommentData `json:"data"`
}

type Comment struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    CommentData `json:"data"`
}
