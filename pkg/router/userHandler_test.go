package router

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
var mockUser dto.User
var userJSON []byte
var userBody *bytes.Buffer
var invalidUser dto.User
var invalidUserJSON []byte
var invalidUserBody *bytes.Buffer

func setupUserTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService = service.NewMockUserService(ctrl)
	uh = UserHandler{mockUserService}
	mockUsers = []dto.User{
		{UUID: uuid.New(), Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: uuid.New(), Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}

	mockUser = mockUsers[0]
	userJSON, _ = json.Marshal(mockUser)
	userBody = bytes.NewBuffer(userJSON)

	invalidUser = dto.User{UUID: uuid.New(), Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"}
	invalidUserJSON, _ = json.Marshal(invalidUser)
	invalidUserBody = bytes.NewBuffer(invalidUserJSON)

	userRouter = mux.NewRouter()
	userRouter.HandleFunc("/user", uh.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/user", uh.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/user/{slug}", uh.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/user/{slug}", uh.RemoveUser).Methods(http.MethodDelete)
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

func Test_all_user_route_should_return_200_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().GetAllUsers().Return(mockUsers, nil)

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}

func Test_user_update_route_returns_500_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().UpdateUser(userJSON).Return(nil, errs.NewUnexpectedError("Couldn't update user"))

	req, _ := http.NewRequest(http.MethodPut, "/user", userBody)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_update_user_route_returns_400_error(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().UpdateUser(invalidUserJSON).Return(nil, errs.NewInvalidError("User name must not be blank"))

	req, _ := http.NewRequest(http.MethodPut, "/user", invalidUserBody)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", recorder.Code)
	}
}

func Test_user_update_route_returns_200_code(t *testing.T) {
	setupUserTest(t)

	mockUserService.EXPECT().UpdateUser(userJSON).Return(&mockUser, nil)

	req, _ := http.NewRequest(http.MethodPut, "/user", userBody)

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

	mockUserService.EXPECT().SingleUser("testuser").Return(&mockUser, nil)

	req, _ := http.NewRequest(http.MethodGet, "/user/testuser", nil)

	recorder := httptest.NewRecorder()
	userRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 500, but got", recorder.Code)
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
