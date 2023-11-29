package router

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ChatRoomHandler struct {
	RoomService    *service.Room
	CommentService service.CommentService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		origin := req.Header.Get("Origin")
		log.Printf("Upgrade request from %s", origin)
		return origin == conf.App.HttpProto+conf.App.Domain
	},
}

func (crh ChatRoomHandler) Connect(resp http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value("auth").(dto.Auth).UserId
	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &service.ChatClient{Room: crh.RoomService, UserId: userId, Conn: conn, Send: make(chan []byte, 256), ComServ: crh.CommentService}
	client.Room.Register <- client

	log.Printf("user %s has connected to the chatroom", userId)
	go client.Pub()
	go client.Sub()
}
