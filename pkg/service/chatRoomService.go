package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	commentRepo domain.CommentRepository
	clients     map[*Client]bool
	broadcast   chan []byte
	Register    chan *Client
	unregister  chan *Client
}

func NewRoom(commentRepositoryDB domain.CommentRepositoryDb) *Room {
	return &Room{
		commentRepo: commentRepositoryDB,
		broadcast:   make(chan []byte),
		Register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.clients[client] = true
			comments, err := r.commentRepo.FindAllComments()
			if err != nil {
				return
			}

			for _, comment := range comments {
				commentJSON, _ := json.Marshal(comment.ToDto())
				client.Send <- commentJSON
			}
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.Send)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				commentDTO := dto.Comment{UUID: uuid.New(), UserId: client.UserId, Text: string(message), CreatedAt: time.Now()}
				commentJSON, err := json.Marshal(commentDTO)
				if err != nil {
					return
				}
				select {
				case client.Send <- commentJSON:
					_, err := r.commentRepo.CreateComment(commentDTO)
					if err != nil {
						return
					}
				default:
					close(client.Send)
					delete(r.clients, client)
				}
			}
		}
	}
}
