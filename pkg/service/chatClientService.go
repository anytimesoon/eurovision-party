package service

import (
	"bytes"
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 5120
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
			log.Printf("Chatroom pub connection closed for %s. %s", c.UserId, err)
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
			log.Printf("Failed to reset read deadline for user %s in pong handler. %s", c.UserId, err)
			return err
		}
		//log.Printf("Received pong from user %s", c.UserId)
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

		// decode message
		filteredMessage := dto.SocketMessage{}
		err = json.Unmarshal(message, &filteredMessage)
		if err != nil {
			log.Println("Failed to unmarshal message")
			continue
		}

		switch filteredMessage.Category {
		case enum.COMMENT:
			commentJSON, appErr := c.ComServ.CreateComment(filteredMessage.Body)
			if appErr != nil {
				return
			}
			log.Println("New message received from", c.UserId.String())
			c.Room.broadcastChatMessage <- commentJSON
		case enum.LATEST_COMMENT:
			commentsJSON, appErr := c.ComServ.FindCommentsAfter(filteredMessage.Body)
			if appErr != nil {
				return
			}
			log.Println("Sending latest messages to", c.UserId.String())
			chatMessages := dto.SocketMessage{
				Category: enum.COMMENT_ARRAY,
				Body:     commentsJSON,
			}
			log.Println(chatMessages)
			chatMessagesJSON, err := json.Marshal(chatMessages)
			if err != nil {
				log.Println("Failed to marshal latest messages for user", c.UserId)
			}

			c.Send <- chatMessagesJSON
		default:
			log.Printf("Message category not recognised")
		}
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
		case message, ok := <-c.Send:
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

			_, err = writer.Write(message)
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
				log.Printf("Failed to set deadline for ping message for user %s. %s", c.UserId, err)
				return
			}
			//log.Printf("Sending ping to user %s", c.UserId)
			err = c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("Failed to send ping message for user %s. %s", c.UserId, err)
				return
			}
		}
	}
}
