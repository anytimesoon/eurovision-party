package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/chatMsgType"
	service2 "github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func TestChatRoomHandler_Connect(t *testing.T) {
	adminSession := getAdminSession()
	regularSession := getRegularSession()

	type fields struct {
		RoomService    *service2.Room
		CommentService service2.CommentService
		AuthService    service2.AuthService
	}

	tests := []struct {
		name          string
		fields        fields
		user          dto2.User
		token         string
		expectUpgrade bool
		expectError   bool
	}{
		{
			name: "successful admin websocket connection",
			fields: fields{
				RoomService:    newTestChatRoomService(),
				CommentService: newTestCommentService(),
				AuthService:    newTestAuthService(),
			},
			user:          adminUserMock,
			token:         adminSession,
			expectUpgrade: true,
			expectError:   false,
		},
		{
			name: "successful regular user websocket connection",
			fields: fields{
				RoomService:    newTestChatRoomService(),
				CommentService: newTestCommentService(),
				AuthService:    newTestAuthService(),
			},
			user:          regularUserMock,
			token:         regularSession,
			expectUpgrade: true,
			expectError:   false,
		},
		{
			name: "unauthorized connection attempt",
			fields: fields{
				RoomService:    newTestChatRoomService(),
				CommentService: newTestCommentService(),
				AuthService:    newTestAuthService(),
			},
			user:          regularUserMock,
			token:         "invalid_token",
			expectUpgrade: false,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Add the token to the request context
				r = mux.SetURLVars(r, map[string]string{"token": tt.token})

				crh := ChatRoomHandler{
					RoomService:    tt.fields.RoomService,
					CommentService: tt.fields.CommentService,
					AuthService:    tt.fields.AuthService,
				}

				go crh.RoomService.Run()
				crh.Connect(w, r)
			}))
			defer server.Close()

			// Convert http://... to ws://...
			url := "ws" + strings.TrimPrefix(server.URL, "http") + "/api/ws"

			// Connect to the server
			ws, resp, err := websocket.DefaultDialer.Dial(url, http.Header{
				"Origin": []string{conf.App.HttpProto + conf.App.Domain},
			})

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if !tt.expectUpgrade {
				if err == nil {
					t.Error("Expected connection to fail but it succeeded")
				}
				return
			}

			if err != nil {
				t.Fatalf("Could not open websocket connection: %v", err)
			}
			defer ws.Close()
			defer resp.Body.Close()

			// Verify the connection was upgraded
			if resp.StatusCode != http.StatusSwitchingProtocols {
				t.Errorf("Expected status code %d but got %d", http.StatusSwitchingProtocols, resp.StatusCode)
			}

			mockComment := dto2.Comment{
				UUID:      uuid.New(),
				UserId:    tt.user.UUID,
				Text:      "hello",
				FileName:  "fabulous.file",
				CreatedAt: time.Now(),
				ReplyTo:   nil,
			}
			comm, err := json.Marshal(mockComment)
			if err != nil {
				panic(err)
			}

			mockMessage := dto2.SocketMessage{
				Category: chatMsgType.COMMENT,
				Body:     comm,
			}

			mockMessageJSON, err := json.Marshal(mockMessage)
			if err != nil {
				panic(err)
			}

			err = ws.WriteMessage(websocket.TextMessage, mockMessageJSON)
			if err != nil {
				t.Fatalf("Could not write message: %v", err)
			}

			// Read response with timeout
			done := make(chan struct{})
			go func() {
				defer close(done)
				_, message, err := ws.ReadMessage()
				if err != nil {
					t.Errorf("Failed to read message: %v", err)
					return
				}
				t.Logf("Received message: %s", message)

				persistedComment := &dao.Comment{}
				testDB.Get(mockComment.UUID.String(), persistedComment)
				if persistedComment.Text != mockComment.Text {
					t.Errorf("Comment text does not match. Expected %s but got %s", mockComment.Text, persistedComment.Text)
				}
				if persistedComment.UserId != mockComment.UserId {
					t.Errorf("Comment user id does not match. Expected %s but got %s", mockComment.UserId, persistedComment.UserId)
				}
				if persistedComment.FileName != mockComment.FileName {
					t.Errorf("Comment file name does not match. Expected %s but got %s", mockComment.FileName, persistedComment.FileName)
				}
				if persistedComment.ReplyTo != nil {
					t.Errorf("Comment reply to does not match. Expected %+v but got %+v", mockComment.ReplyTo, persistedComment.ReplyTo)
				}
			}()

			select {
			case <-done:

			case <-time.After(5 * time.Second):
				t.Fatal("Test timed out waiting for response")
			}

		})
	}
}
