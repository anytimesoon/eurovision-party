package dto

import "encoding/json"

type ChatMessage struct {
	Category string          `json:"category"`
	Body     json.RawMessage `json:"body"`
}
