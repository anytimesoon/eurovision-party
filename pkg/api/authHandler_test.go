package api

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/timshannon/bolthold"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthHandler_Login(t *testing.T) {
	type (
		fields struct {
			Service service.AuthService
		}
		args struct {
			resp http.ResponseWriter
			req  *http.Request
		}
		expected struct {
			statusCode int
			hasSession bool
		}
	)
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful login",
			fields: fields{
				Service: newTestAuthService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(fmt.Sprintf(`{
                    "token": "adminToken",
                    "userId": "%s"
                }`, adminUserId.String()))),
			},
			expected: expected{
				statusCode: http.StatusOK,
				hasSession: true,
			},
		},
		{
			name: "failed login - no user",
			fields: fields{
				Service: newTestAuthService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(fmt.Sprintf(`{
                    "token": "badToken",
                    "userId": ""
                }`))),
			},
			expected: expected{
				statusCode: http.StatusInternalServerError,
				hasSession: false,
			},
		},
		{
			name: "failed login - no token",
			fields: fields{
				Service: newTestAuthService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(fmt.Sprintf(`{
                    "token": "",
                    "userId": "%s"
                }`, adminUserId.String()))),
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				hasSession: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := AuthHandler{
				Service: tt.fields.Service,
			}
			ah.Login(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("Login return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			count, err := testDB.Count(&dao.Session{}, &bolthold.Query{})
			if err != nil {
				panic(err)
			}
			if count > 0 && !tt.expected.hasSession {
				removeAllSessions()
				t.Errorf("Login has sessions = %v, want %v", count, 0)
			}
			if count != 1 && tt.expected.hasSession {
				removeAllSessions()
				t.Errorf("Login has sessions = %v, want %v", count, 1)
			}
			removeAllSessions()
		})
	}
}

func removeAllSessions() {
	err := testDB.DeleteMatching(&dao.Session{}, &bolthold.Query{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to delete sessions from test db!")
	}
	afterCount, _ := testDB.Count(&dao.Session{}, &bolthold.Query{})
	fmt.Printf("Deleted sessions from test db! %d left", afterCount)
}
