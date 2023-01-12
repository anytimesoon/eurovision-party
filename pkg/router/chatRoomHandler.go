package router

import (
	"eurovision/pkg/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ChatRoomHandler struct {
	AuthService    service.AuthService
	RoomService    *service.Room
	CommentService service.CommentService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		origin := req.Header.Get("Origin")
		log.Printf("Upgrade request from %s", origin)
		//return origin == "http://127.0.0.1:3000"
		return true
	},
}

func (crh ChatRoomHandler) Connect(resp http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Sec-WebSocket-Protocol")

	user, appErr := authService.AuthorizeChat(token)
	if appErr != nil {
		log.Printf("%s method %s was requested by %q and rejected because token was rejected. %s", req.RequestURI, req.Method, req.RemoteAddr, appErr)
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	var header http.Header

	conn, err := upgrader.Upgrade(resp, req, header)
	if err != nil {
		log.Println(err)
		return
	}
	client := &service.ChatClient{Room: crh.RoomService, UserId: user.UUID, Conn: conn, Send: make(chan []byte, 256), ComServ: crh.CommentService}
	client.Room.Register <- client

	go client.Pub()
	go client.Sub()
}
