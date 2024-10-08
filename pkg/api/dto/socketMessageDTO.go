package dto

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
)

type SocketMessage struct {
	Category enum.ChatMsgType `json:"category"`
	Body     json.RawMessage  `json:"body"`
}

type SocketError string

type ChatResponseBody interface {
	*Comment | []Comment | SocketError | UpdateMessage | *NewUser
}

// NewSocketErrorMessage can be passed an empty string to get the default message
func NewSocketErrorMessage(message string) SocketMessage {
	if message == "" {
		message = "Some pyro's went off back stage. We lost your message in the distraction 😟"
	}
	encodedString, _ := json.Marshal(message)
	return SocketMessage{
		Category: enum.ERROR,
		Body:     encodedString,
	}
}

func NewSocketMessage[T ChatResponseBody](category enum.ChatMsgType, payload T) SocketMessage {
	body, err := json.Marshal(payload)
	if err != nil {
		return NewSocketErrorMessage("")
	}

	return SocketMessage{
		Category: category,
		Body:     body,
	}
}
