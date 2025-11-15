package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	dao2 "github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	service2 "github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/timshannon/bolthold"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	type fields struct {
		Service      service2.UserService
		AssetService service2.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		users      map[uuid.UUID]dto2.User
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
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/api/user", nil),
			},
			expected: expected{
				statusCode: http.StatusOK,
				users: map[uuid.UUID]dto2.User{
					regularUserId:        regularUserMock,
					adminUserId:          adminUserMock,
					friendOfFriendUserId: friendOfFriendUserMock.ToDto(),
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

			var usersDto dto2.ApiPayload[map[uuid.UUID]dto2.User]
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
		Service      service2.UserService
		AssetService service2.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		users      dto2.User
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
				AssetService: service2.NewAssetService(),
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

			var resultDto dto2.ApiPayload[*dto2.User]
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
		Service      service2.UserService
		AssetService service2.AssetService
	}
	type args struct {
		resp             http.ResponseWriter
		req              *http.Request
		requestingUserId uuid.UUID
	}
	type expected struct {
		statusCode int
		users      []*dto2.NewUser
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful get registered users as admin",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp:             httptest.NewRecorder(),
				req:              httptest.NewRequest(http.MethodGet, "/api/user/registered", nil),
				requestingUserId: adminUserId,
			},
			expected: expected{
				statusCode: http.StatusOK,
				users:      generateNewUsers(),
			},
		},
		{
			name: "successful get registered users as user",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp:             httptest.NewRecorder(),
				req:              httptest.NewRequest(http.MethodGet, "/api/user/registered", nil),
				requestingUserId: regularUserId,
			},
			expected: expected{
				statusCode: http.StatusOK,
				users:      newUsersFilteredById(regularUserId),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := UserHandler{
				UserService:  tt.fields.Service,
				AssetService: tt.fields.AssetService,
			}
			tt.args.req = mux.SetURLVars(tt.args.req, map[string]string{"userId": tt.args.requestingUserId.String()})
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

			var usersDto dto2.ApiPayload[[]*dto2.NewUser]
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

func TestUserHandler_Register_LimitsTheAmountOfUserRegistrations(t *testing.T) {
	userService := newTestUserService()
	handler := UserHandler{
		UserService:  userService,
		AssetService: service2.NewAssetService(),
	}

	// Test regular user hitting the limit
	regularUserTests := []struct {
		name          string
		expectedCode  int
		expectedError string
		registrations int
	}{
		{
			name:          "can register users before hitting limit",
			expectedCode:  http.StatusOK,
			registrations: 3,
		},
		{
			name:          "can register last user at limit",
			expectedCode:  http.StatusOK,
			registrations: 1,
		},
		{
			name:          "cannot register users after hitting limit",
			expectedCode:  http.StatusForbidden,
			expectedError: errs.Common.MaxInvitesExceeded,
			registrations: 4,
		},
	}

	for _, tt := range regularUserTests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.registrations; i++ {
				newUser := dto2.NewUser{
					Name:      fmt.Sprintf("Test User %d", i),
					CreatedBy: regularUserId,
				}
				body, _ := json.Marshal(newUser)
				req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
				rec := httptest.NewRecorder()

				handler.Register(rec, req)

				result := rec.Result()
				defer result.Body.Close()

				if i == tt.registrations && result.StatusCode != tt.expectedCode {
					t.Errorf("expected status code %d, got %d", tt.expectedCode, result.StatusCode)
				}

				if i == tt.registrations && tt.expectedError != "" {
					var response dto2.ApiPayload[*dto2.NewUser]
					if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
						t.Fatal(err)
					}
					if response.Error != tt.expectedError {
						t.Errorf("expected error message %q, got %q", tt.expectedError, response.Error)
					}
				}
			}
		})
	}

	// Test admin user not being affected by the limit
	t.Run("admin can register more than MAX_INVITES users", func(t *testing.T) {
		for i := 0; i < 7; i++ { // Testing with more than MAX_INVITES (5)
			newUser := dto2.NewUser{
				Name:      fmt.Sprintf("Admin Created User %d", i),
				CreatedBy: adminUserId,
			}
			body, _ := json.Marshal(newUser)
			req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			handler.Register(rec, req)

			result := rec.Result()
			defer result.Body.Close()

			if result.StatusCode != http.StatusOK {
				t.Errorf("admin failed to create user %d, got status code %d", i, result.StatusCode)
			}
		}
	})
}

func TestUserHandler_Register(t *testing.T) {
	type fields struct {
		Service      service2.UserService
		AssetService service2.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		user       *dto2.NewUser
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "successful register user by admin",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto2.NewUser{
						Name:      "Test User",
						CreatedBy: adminUserId,
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto2.NewUser{
					Name:      "Test User",
					Slug:      "test-user",
					CreatedBy: adminUserId,
					AuthLvl:   authLvl.USER,
				},
			},
		},
		{
			name: "unauthorized register attempt",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto2.NewUser{
						Name:      "Test User",
						CreatedBy: friendOfFriendUserId,
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.USER})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusForbidden,
				user:       nil,
			},
		},
		{
			name: "successful unique slug",
			fields: fields{
				Service:      newTestUserService(),
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					newUser := &dto2.NewUser{
						Name:      regularUserMock.Name,
						CreatedBy: adminUserId,
					}
					body, err := json.Marshal(newUser)
					if err != nil {
						t.Fatal(err)
					}
					req := httptest.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(body))
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto2.NewUser{
					Name:      regularUserMock.Name,
					Slug:      regularUserMock.Slug + "-1",
					CreatedBy: adminUserId,
					AuthLvl:   authLvl.USER,
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

				var userDto dto2.ApiPayload[*dto2.NewUser]
				err = json.Unmarshal(resultBody, &userDto)
				if err != nil {
					panic(err)
				}

				if userDto.Body.Name != tt.expected.user.Name ||
					userDto.Body.Slug != tt.expected.user.Slug ||
					userDto.Body.AuthLvl != tt.expected.user.AuthLvl ||
					userDto.Body.CreatedBy != tt.expected.user.CreatedBy {
					t.Errorf("Registered user doesn't match:\ngot:  %+v\nwant: %+v",
						userDto.Body, tt.expected.user)
				}

				countryCount, err := testDB.Count(dao2.Country{}, &bolthold.Query{})
				if err != nil {
					panic(err)
				}

				var votes []dao2.Vote
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
		Service      service2.UserService
		AssetService service2.AssetService
	}
	type args struct {
		resp http.ResponseWriter
		req  *http.Request
	}
	type expected struct {
		statusCode int
		user       *dto2.User
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
				AssetService: service2.NewAssetService(),
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
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.ADMIN})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto2.User{
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
				AssetService: service2.NewAssetService(),
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
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{
						UserId:  regularUserMock.UUID,
						AuthLvl: authLvl.USER,
					})
					return req.WithContext(ctx)
				}(),
			},
			expected: expected{
				statusCode: http.StatusOK,
				user: &dto2.User{
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
				AssetService: service2.NewAssetService(),
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
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{
						UserId:  uuid.New(), // Different user ID
						AuthLvl: authLvl.USER,
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

				var userDto dto2.ApiPayload[*dto2.User]
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
		Service      service2.UserService
		AssetService service2.AssetService
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
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/api/user/regular", nil)
					req = mux.SetURLVars(req, map[string]string{"slug": regularUserMock.Slug})
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.ADMIN})
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
				AssetService: service2.NewAssetService(),
			},
			args: args{
				resp: httptest.NewRecorder(),
				req: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/api/user/regular", nil)
					req = mux.SetURLVars(req, map[string]string{"slug": regularUserMock.Slug})
					ctx := context.WithValue(req.Context(), "auth", dto2.Auth{AuthLvl: authLvl.USER})
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
