package api

import (
	"log"
	"net/http"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ChatRoomHandler struct {
	RoomService    *service.Room
	CommentService service.CommentService
	AuthService    service.AuthService
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
	params := mux.Vars(req)

	token, appErr := crh.AuthService.Authorize(params["token"])
	if appErr != nil {
		log.Println(appErr)
		return
	}

	log.Printf("Connecting user %s to chat", token.UserId)

	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &service.ChatClient{Room: crh.RoomService, UserId: token.UserId, Conn: conn, Send: make(chan []byte, 256), ComServ: crh.CommentService}
	client.Room.Register <- client

	log.Printf("user %s has connected to the chatroom", token.UserId)
	go client.Pub()
	go client.Sub()
}
