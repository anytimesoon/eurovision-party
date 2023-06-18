package service

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"github.com/google/uuid"
	"log"
)

type Room struct {
	CommentService CommentService
	clients        map[uuid.UUID]*ChatClient
	broadcast      chan []byte
	Register       chan *ChatClient
	unregister     chan *ChatClient
}

func NewRoom(commentService CommentService) *Room {
	return &Room{
		CommentService: commentService,
		broadcast:      make(chan []byte),
		Register:       make(chan *ChatClient),
		unregister:     make(chan *ChatClient),
		clients:        make(map[uuid.UUID]*ChatClient),
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
			chatMessages := dto.ChatMessage{
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
		case commentJSON := <-r.broadcast:
			chatMessage := dto.ChatMessage{
				Category: enum.COMMENT,
				Body:     commentJSON,
			}

			message, err := json.Marshal(chatMessage)
			if err != nil {
				log.Printf("failed to encode message to chatMessage")
				break
			}
			for userId, client := range r.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.clients, userId)
				}
			}
		}
	}
}
