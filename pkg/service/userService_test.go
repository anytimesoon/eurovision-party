package service

import (
	"eurovision/mocks/domain"
	realDomain "eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var userService UserService
var mockUserRepository *domain.MockUserRepository
var mockUsers []realDomain.User
var mockUser realDomain.User
var mockUsersDTO []dto.User
var mockUserDTO dto.User

func setupUserTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository = domain.NewMockUserRepository(ctrl)
	userService = DefaultUserService{mockUserRepository}
	mockUsers = []realDomain.User{
		{UUID: uuid.New(), AuthLvl: realDomain.Admin, Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: uuid.New(), AuthLvl: realDomain.None, Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}
	mockUser = mockUsers[0]
	mockUsersDTO = []dto.User{
		{UUID: mockUsers[0].UUID, Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: mockUsers[1].UUID, Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}
	mockUserDTO = mockUsersDTO[0]
}

func Test_user_service_returns_all_users(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().FindAllUsers().Return(mockUsers, nil)

	result, _ := userService.GetAllUsers()

	if result[0] != mockUsersDTO[0] || result[1] != mockUsersDTO[1] {
		t.Error("Returned users do not match expected")
	}
}

func Test_all_user_service_returns_error(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().FindAllUsers().Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := userService.GetAllUsers()

	if err == nil {
		t.Error("Was expecting and error, but got none")
	}
}
