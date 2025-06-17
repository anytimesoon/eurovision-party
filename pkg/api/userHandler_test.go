package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/timshannon/bolthold"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		users      map[uuid.UUID]dto.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get all users",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/user", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				users: map[uuid.UUID]dto.User{
					regularUserMock.UUID: regularUserMock,
					adminUserMock.UUID:   adminUserMock,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			uh.GetAllUsers(tt.args.resp, tt.args.req)
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

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var usersDto dto.ApiPayload[map[uuid.UUID]dto.User]
			err = json.Unmarshal(resultBody, &usersDto)
			if err != nil {
				panic(err)
			}

			userCount := len(usersDto.Body)
			if userCount != len(tt.expected.users) {
				t.Errorf("User count = %v, want %v", userCount, len(tt.expected.users))
			}

			for _, expectedUser := range tt.expected.users {
				found := false
				for _, actualUser := range usersDto.Body {
					if expectedUser.UUID == actualUser.UUID {
						if expectedUser != actualUser {
							t.Errorf("User with ID %v doesn't match:\ngot:  %+v\nwant: %+v",
								actualUser.UUID, actualUser, expectedUser)
						}
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected user not found: %+v", expectedUser)
				}
			}

		})
	}
}

func TestUserHandler_GetOneUser(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		users      dto.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get one user",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/user/regular", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				users:      regularUserMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"slug": regularUserMock.Slug})
			uh.GetOneUser(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("GetAllCountries return status code = %v, want %v", result.StatusCode, http.StatusOK)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var resultDto dto.ApiPayload[*dto.User]
			err = json.Unmarshal(resultBody, &resultDto)
			if err != nil {
				panic(err)
			}

			if resultDto.Body.UUID != tt.expected.users.UUID {
				t.Errorf("User %s not found in response", tt.expected.users.UUID)
			}
		})
	}
}

func TestUserHandler_GetRegisteredUsers(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		users      []*dto.NewUser
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get registered users",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/user/registered", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				users:      generateNewUsers(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			uh.GetRegisteredUsers(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("GetRegisteredUsers return status code = %v, want %v", result.StatusCode, tt.expected.statusCode)
			}

			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				panic(err)
			}

			var usersDto dto.ApiPayload[[]*dto.NewUser]
			err = json.Unmarshal(resultBody, &usersDto)
			if err != nil {
				panic(err)
			}

			userCount := len(usersDto.Body)
			if userCount != len(tt.expected.users) {
				t.Errorf("User count = %v, want %v", userCount, len(tt.expected.users))
			}

			for _, expectedUser := range tt.expected.users {
				found := false
				for _, actualUser := range usersDto.Body {
					if expectedUser.UUID == actualUser.UUID {
						if expectedUser.Name != actualUser.Name ||
							expectedUser.Slug != actualUser.Slug ||
							expectedUser.AuthLvl != actualUser.AuthLvl {
							t.Errorf("User with UUID %v doesn't match:\ngot:  %+v\nwant: %+v",
								expectedUser.UUID, actualUser, expectedUser)
						}
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected user not found: %+v", expectedUser)
				}

			}
		})
	}
}

func TestUserHandler_Register(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		user       *dto.NewUser
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful register user",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto.NewUser{
						Name: "Test User",
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto.NewUser{
					Name:    "Test User",
					Slug:    "test-user",
					AuthLvl: 0,
				},
			},
		},
		{
			name: "unauthorized register attempt",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto.NewUser{
						Name: "Test User",
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.NONE})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				user:       nil,
			},
		},
		{
			name: "successful unique slug",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto.NewUser{
						Name: regularUserMock.Name,
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto.NewUser{
					Name:    regularUserMock.Name,
					Slug:    regularUserMock.Slug + "-1",
					AuthLvl: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			uh.Register(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("Register return status code = %v, want %v", result.StatusCode, tt.expected.statusCode)
			}

			if tt.expected.user != nil {
				resultBody, err := io.ReadAll(result.Body)
				if err != nil {
					panic(err)
				}

				var userDto dto.ApiPayload[*dto.NewUser]
				err = json.Unmarshal(resultBody, &userDto)
				if err != nil {
					panic(err)
				}

				if userDto.Body.Name != tt.expected.user.Name ||
					userDto.Body.Slug != tt.expected.user.Slug {
					t.Errorf("Registered user doesn't match:\ngot:  %+v\nwant: %+v",
						userDto.Body, tt.expected.user)
				}

				countryCount, err := testDB.Count(dao.Country{}, &bolthold.Query{})
				if err != nil {
					panic(err)
				}

				var votes []dao.Vote
				err = testDB.Find(&votes, &bolthold.Query{})
				if err != nil {
					panic(err)
				}

				voteCount := 0
				for _, vote := range votes {
					if vote.UserId == userDto.Body.UUID {
						voteCount++
					}
				}

				if countryCount != voteCount {
					t.Errorf("Vote count doesn't match:\ngot:  %v\nwant: %v",
						voteCount, countryCount)
				}
			}
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		user       *dto.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful update user by admin",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					updatedUser := regularUserMock
					updatedUser.Name = "Updated Name"
					body, err := json.Marshal(updatedUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPut, "/api/user", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto.User{
					UUID: regularUserMock.UUID,
					Name: "Updated Name",
					Slug: regularUserMock.Slug,
					Icon: regularUserMock.Icon,
				},
			},
		},
		{
			name: "successful update own user",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					updatedUser := regularUserMock
					updatedUser.Name = "Updated Name"
					body, err := json.Marshal(updatedUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPut, "/api/user", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{
						UserId:  regularUserMock.UUID,
						AuthLvl: enum.NONE,
					})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto.User{
					UUID: regularUserMock.UUID,
					Name: "Updated Name",
					Slug: regularUserMock.Slug,
					Icon: regularUserMock.Icon,
				},
			},
		},
		{
			name: "unauthorized update attempt",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					updatedUser := regularUserMock
					updatedUser.Name = "Updated Name"
					body, err := json.Marshal(updatedUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPut, "/api/user", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{
						UserId:  uuid.New(), // Different user ID
						AuthLvl: enum.NONE,
					})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				user:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			uh.UpdateUser(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("UpdateUser return status code = %v, want %v", result.StatusCode, tt.expected.statusCode)
				return
			}

			if tt.expected.user != nil {
				resultBody, err := io.ReadAll(result.Body)
				if err != nil {
					panic(err)
				}

				var userDto dto.ApiPayload[*dto.User]
				err = json.Unmarshal(resultBody, &userDto)
				if err != nil {
					panic(err)
				}

				if userDto.Body.Name != tt.expected.user.Name ||
					userDto.Body.UUID != tt.expected.user.UUID ||
					userDto.Body.Slug != tt.expected.user.Slug {
					t.Errorf("Updated user doesn't match:\ngot:  %+v\nwant: %+v",
						userDto.Body, tt.expected.user)
				}
			}
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	type fields struct {
		Service      service.UserService
		AssetService service.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful delete user",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/api/user/regular", nil)
					req = mux.SetURLVars(req, map[string]string{"slug": regularUserMock.Slug})
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "unauthorized delete attempt",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/api/user/regular", nil)
					req = mux.SetURLVars(req, map[string]string{"slug": regularUserMock.Slug})
					ctx := context.WithValue(req.Context(), "auth", dto.Auth{AuthLvl: enum.NONE})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			uh.DeleteUser(tt.args.resp, tt.args.req)
			result := tt.args.resp.(*httptest.ResponseRecorder).Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(result.Body)

			if result.StatusCode != tt.expected.statusCode {
				t.Errorf("DeleteUser return status code = %v, want %v", result.StatusCode, tt.expected.statusCode)
			}
		})
	}
}
