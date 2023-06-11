package service

import (
	"encoding/json"
	"eurovision/pkg/dto"
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
			comments, err := r.CommentService.FindAllComments()
			if err != nil {
				return
			}
			log.Printf("%+v", r.clients)
			for _, comment := range comments {
				commentJSON, err := json.Marshal(comment)
				if err != nil {
					log.Printf("Something went wrong when sending a message during registration for user %s. %s", client.UserId, err)
				}
				client.Send <- commentJSON
			}
		case client := <-r.unregister:
			if _, ok := r.clients[client.UserId]; ok {
				delete(r.clients, client.UserId)
				close(client.Send)
			}
		case commentJSON := <-r.broadcast:
			chatMessage := dto.ChatMessage{
				Category: "comment",
				Body:     commentJSON,
			}

			message, err := json.Marshal(chatMessage)
			if err != nil {
				log.Printf("failed to encode message to chatMessage")
				break
			}
			for userId, client := range r.clients {
				//select {
				//case
				log.Println("sending to", userId)
				client.Send <- message
				//default:
				//	close(client.Send)
				//	delete(r.clients, client)
				//}
			}
		}
	}
}
