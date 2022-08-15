package handler

import (
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

var router *mux.Router
var uh UserHandler
var mockService *service.MockUserService

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockUserService(ctrl)
	uh = UserHandler{mockService}

	router = mux.NewRouter()
	router.HandleFunc("/user", uh.FindAllUsers)
}

func Test_should_return_users_with_200_code(t *testing.T) {
	setup(t)
	mockUsers := []dto.User{
		{UUID: uuid.New(), Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: uuid.New(), Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}
	mockService.EXPECT().GetAllUsers().Return(mockUsers, nil)

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_500_code(t *testing.T) {
	setup(t)

	mockService.EXPECT().GetAllUsers().Return(nil, errs.NewUnexpectedError("Couldn't find users"))

	req, _ := http.NewRequest(http.MethodGet, "/user", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
