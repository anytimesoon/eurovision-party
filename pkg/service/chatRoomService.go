package service

import (
	"encoding/json"
)

type Room struct {
	CommentService CommentService
	clients        map[*ChatClient]bool
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
		clients:        make(map[*ChatClient]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.clients[client] = true
			comments, err := r.CommentService.FindAllComments()
			if err != nil {
				return
			}

			for _, comment := range comments {
				commentJSON, _ := json.Marshal(comment)
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
