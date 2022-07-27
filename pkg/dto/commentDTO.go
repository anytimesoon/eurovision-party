package dto

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
