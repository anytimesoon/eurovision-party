package service

import (
	"bytes"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type ChatClient struct {
	Room    *Room
	UserId  uuid.UUID
	Conn    *websocket.Conn
	Send    chan []byte
	ComServ CommentService
}

func (c *ChatClient) Pub() {
	defer func() {
		c.Room.unregister <- c
		err := c.Conn.Close()
		if err != nil {
			log.Printf("Chatroom pub connection closed unexpectedly for %s., %s", c.UserId, err)
			return
		}
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Printf("Failed to set read deadline for user %s. %s", c.UserId, err)
		return
	}
	c.Conn.SetPongHandler(func(string) error {
		err = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Printf("Failed to set read deadline for user %s in pong handler. %s", c.UserId, err)
			return err
		}
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		commentJSON, appErr := c.ComServ.CreateComment(message, c.UserId)
		if appErr != nil {
			return
		}

		c.Room.broadcast <- commentJSON
	}
}

func (c *ChatClient) Sub() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.Conn.Close()
		if err != nil {
			log.Printf("Chat client sub connection closed unexpectedly for %s., %s", c.UserId, err)
			return
		}
	}()
	for {
		select {
		case commentJSON, ok := <-c.Send:
			err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Printf("Failed to set write deadline for user %s. %s", c.UserId, err)
				return
			}
			if !ok {
				// The chatroom closed the channel.
				err = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Printf("Failed to write close message for user %s. %s", c.UserId, err)
					return
				}
				return
			}

			writer, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = writer.Write(commentJSON)
			if err != nil {
				log.Printf("Failed to write chat message for user %s. %s", c.UserId, err)
				return
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, err = writer.Write(newline)
				if err != nil {
					log.Printf("Failed to write new line message for user %s. %s", c.UserId, err)
					return
				}
				_, err = writer.Write(<-c.Send)
				if err != nil {
					log.Printf("Failed to send message for user %s. %s", c.UserId, err)
					return
				}
			}

			err = writer.Close()
			if err != nil {
				log.Printf("Chatclient writer closed unexpectedly for user %s. %s", c.UserId, err)
				return
			}
		case <-ticker.C:
			err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Printf("Failed to write new line message for user %s. %s", c.UserId, err)
				return
			}
			err = c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("Failed to write ping message for user %s. %s", c.UserId, err)
				return
			}
		}
	}
}
