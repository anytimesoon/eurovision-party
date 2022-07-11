package dto

import "github.com/google/uuid"

type CommentData struct {
	UUID   uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Text   string    `json:"text"`
}

type Comments struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []CommentData `json:"data"`
}
