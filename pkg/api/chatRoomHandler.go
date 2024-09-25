package api

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	log.Printf("Connecting user %s to chat", params["u"])
	appErr := crh.AuthService.AuthorizeChat(params["t"], params["u"])
	if appErr != nil {
		log.Println(appErr)
		return
	}

	// not processing the error because if we get here, we know the uuid is valid
	userId, _ := uuid.Parse(params["u"])
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
