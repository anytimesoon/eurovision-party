package dto

import (
	"encoding/json"
	"eurovision/pkg/enum"
)

type ChatMessage struct {
	Category enum.ChatMsgType `json:"category"`
	Body     json.RawMessage  `json:"body"`
}
