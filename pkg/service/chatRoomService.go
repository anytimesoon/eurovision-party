package service

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/google/uuid"
	"log"
)

type Room struct {
	CommentService       CommentService
	clients              map[uuid.UUID]*ChatClient
	broadcastChatMessage chan []byte
	BroadcastUpdate      chan []byte
	Register             chan *ChatClient
	unregister           chan *ChatClient
}

func NewRoom(commentService CommentService) *Room {
	return &Room{
		CommentService:       commentService,
		broadcastChatMessage: make(chan []byte),
		BroadcastUpdate:      make(chan []byte),
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
			comments, appErr := r.CommentService.FindAllComments()
			if appErr != nil {
				return
			}
			log.Printf("%+v", r.clients)
			commentsJSON, err := json.Marshal(comments)
			if err != nil {
				//TODO: handle error
			}
			chatMessages := dto.SocketMessage{
				Category: enum.COMMENT_ARRAY,
				Body:     commentsJSON,
			}

			chatMessagesJSON, err := json.Marshal(chatMessages)
			if err != nil {
				//TODO: handle error
			}

			client.Send <- chatMessagesJSON

		case client := <-r.unregister:
			if _, ok := r.clients[client.UserId]; ok {
				delete(r.clients, client.UserId)
				close(client.Send)
			}
		case commentJSON := <-r.broadcastChatMessage:
			chatMessage := dto.SocketMessage{
				Category: enum.COMMENT,
				Body:     commentJSON,
			}

			r.broadcast(chatMessage)
		case updateJSON := <-r.BroadcastUpdate:
			chatMessage := dto.SocketMessage{
				Category: enum.UPDATE_USER,
				Body:     updateJSON,
			}

			log.Println("Broadcasting user update")
			r.broadcast(chatMessage)
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
