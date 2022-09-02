package handler

import (
	"bytes"
	"encoding/json"
	"eurovision/mocks/service"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var userRouter *mux.Router
var uh UserHandler
var mockUserService *service.MockUserService
var mockUsers []dto.User

func setupUserTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService = service.NewMockUserService(ctrl)
	uh = UserHandler{mockUserService}
	mockUsers = []dto.User{
		{UUID: uuid.New(), Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: uuid.New(), Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}

	userRouter = mux.NewRouter()
	userRouter.HandleFunc("/user", uh.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/user", uh.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/user/{slug}", uh.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/user/{slug}", uh.RemoveUser).Methods(http.MethodDelete)
}

func Test_all_user_route_should_return_users_with_200_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().GetAllUsers().Return(mockUsers, nil)

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	users := make([]dto.User, 0)
	json.Unmarshal(recorder.Body.Bytes(), &users)

	if len(users) != 2 {
		t.Error("Expected 2 users, but found", len(users))
	}
}

func Test_all_user_route_should_return_500_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().GetAllUsers().Return(nil, errs.NewUnexpectedError("Couldn't find users"))

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_user_update_route_returns_500_code(t *testing.T) {
	setupUserTest(t)

	mockUser := mockUsers[0]
	userJSON, _ := json.Marshal(mockUser)
	body := bytes.NewBuffer(userJSON)

	mockUserService.EXPECT().UpdateUser(userJSON).Return(nil, errs.NewUnexpectedError("Couldn't update user"))

	req, _ := http.NewRequest(http.MethodPut, "/user", body)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_user_update_route_returns_updated_user(t *testing.T) {
	setupUserTest(t)

	mockUser := mockUsers[0]
	userJSON, _ := json.Marshal(mockUser)
	body := bytes.NewBuffer(userJSON)

	mockUserService.EXPECT().UpdateUser(userJSON).Return(&mockUser, nil)

	req, _ := http.NewRequest(http.MethodPut, "/user", body)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}

	var user dto.User
	json.Unmarshal(recorder.Body.Bytes(), &user)

	if user != mockUser {
		t.Errorf("Expected %+v to equal %+v", user, mockUser)
	}
}

func Test_new_user_route_returns_500_code(t *testing.T) {
	setupUserTest(t)

	mockUser := mockUsers[0]
	userJSON, _ := json.Marshal(mockUser)
	body := bytes.NewBuffer(userJSON)

	mockUserService.EXPECT().CreateUser(userJSON).Return(nil, errs.NewUnexpectedError("Couldn't create new user"))

	req, _ := http.NewRequest(http.MethodPost, "/user", body)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_new_user_route_returns_200_code(t *testing.T) {
	setupUserTest(t)

	resultUser := dto.User{UUID: uuid.New(), Name: "newUser", Slug: "newuser", Icon: "/img/static/img/newuser.png"}
	mockUser := dto.User{Name: "newUser"}
	userJSON, _ := json.Marshal(mockUser)
	body := bytes.NewBuffer(userJSON)

	mockUserService.EXPECT().CreateUser(userJSON).Return(&resultUser, nil)

	req, _ := http.NewRequest(http.MethodPost, "/user", body)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}

func Test_find_one_user_route_returns_500_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().SingleUser("testuser").Return(nil, errs.NewUnexpectedError("Couldn't find user"))

	req, _ := http.NewRequest(http.MethodGet, "/user/testuser", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}

func Test_find_one_user_route_returns_user(t *testing.T) {
	setupUserTest(t)

	mockUser := mockUsers[0]

	mockUserService.EXPECT().SingleUser("testuser").Return(&mockUser, nil)

	req, _ := http.NewRequest(http.MethodGet, "/user/testuser", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 500, but got", recorder.Code)
	}

	var returnedUser dto.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &returnedUser)

	if returnedUser != mockUser {
		t.Errorf("Expected %+v to equal %+v", returnedUser, mockUser)
	}
}

func Test_delete_user_route_returns_500_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().DeleteUser("testuser").Return(errs.NewUnexpectedError("Couldn't delete user"))

	req, _ := http.NewRequest(http.MethodDelete, "/user/testuser", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_delete_user_route_returns_200_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().DeleteUser("testuser").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/user/testuser", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}
