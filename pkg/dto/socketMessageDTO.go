package dto

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
)

type SocketMessage struct {
	Category enum.ChatMsgType `json:"category"`
	Body     json.RawMessage  `json:"body"`
}
