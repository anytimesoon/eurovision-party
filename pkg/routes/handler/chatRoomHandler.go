package handler

import (
	"eurovision/pkg/service"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatRoomHandler struct {
	Room *service.Room
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		// origin := req.Header.Get("Origin")
		// return origin == "http://127.0.0.1:3000"
		return true
	},
}

func (crh ChatRoomHandler) Connect(resp http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &service.Client{Room: crh.Room, UserId: uuid.New(), Conn: conn, Send: make(chan []byte, 256), Db: crh.Room.CommentRepo}
	client.Room.Register <- client

	go client.Pub()
	go client.Sub()
}
