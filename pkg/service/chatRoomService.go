package service

import (
	"encoding/json"
	"log"

	"github.com/anytimesoon/eurovision-party/pkg/enum/chatMsgType"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type Room struct {
	clients              map[uuid.UUID]*ChatClient
	broadcastChatMessage chan *dto.Comment
	BroadcastUpdate      chan dto.SocketMessage
	Register             chan *ChatClient
	unregister           chan *ChatClient
}

func NewRoom() *Room {
	return &Room{
		broadcastChatMessage: make(chan *dto.Comment),
		BroadcastUpdate:      make(chan dto.SocketMessage),
		Register:             make(chan *ChatClient),
		unregister:           make(chan *ChatClient),
		clients:              make(map[uuid.UUID]*ChatClient),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.clients[client.UserId] = client

		case client := <-r.unregister:
			if _, ok := r.clients[client.UserId]; ok {
				delete(r.clients, client.UserId)
				close(client.Send)
			}
		case commentJSON := <-r.broadcastChatMessage:
			chatMessage := dto.NewSocketMessage(
				chatMsgType.COMMENT,
				commentJSON,
			)
			r.broadcast(chatMessage)
		case updateMessage := <-r.BroadcastUpdate:
			if updateMessage.Category == chatMsgType.VOTE_NOTIFICATION {
				log.Println("Broadcasting vote update")
				r.broadcast(updateMessage)
			} else {
				log.Println("Broadcasting user update")
				r.broadcast(updateMessage)
			}
		}
	}
}

func (r *Room) broadcast(msg dto.SocketMessage) {
	message, err := json.Marshal(msg)
	if err != nil {
		log.Printf("failed to encode message to chatMessage")
		return
	}
	for _, client := range r.clients {
		client.Send <- message
	}
}
