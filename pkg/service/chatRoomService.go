package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
)

type Room struct {
	CommentRepo domain.CommentRepository
	clients     map[*Client]bool
	broadcast   chan []byte
	Register    chan *Client
	unregister  chan *Client
}

func NewRoom(commentRepositoryDB domain.CommentRepositoryDb) *Room {
	return &Room{
		CommentRepo: commentRepositoryDB,
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
			comments, err := r.CommentRepo.FindAllComments()
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
		case commentJSON := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.Send <- commentJSON:
				default:
					close(client.Send)
					delete(r.clients, client)
				}
			}
		}
	}
}
