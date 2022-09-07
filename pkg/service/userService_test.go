package service

import (
	"encoding/json"
	mockDomain "eurovision/mocks/domain"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var userService UserService
var mockUserRepository *mockDomain.MockUserRepository
var mockUsers []domain.User
var mockUser domain.User
var mockUsersDTO []dto.User
var mockUserDTO dto.User
var userJSON []byte
var invalidUserDTO dto.User
var invalidUserJSON []byte

func setupUserTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository = mockDomain.NewMockUserRepository(ctrl)
	userService = DefaultUserService{mockUserRepository}
	mockUsers = []domain.User{
		{UUID: uuid.New(), AuthLvl: domain.Admin, Name: "tEsTuSeR", Slug: "testuser", Icon: "/img/static/img/newuser.png"},
		{UUID: uuid.New(), AuthLvl: domain.None, Name: "mOcKuSeR", Slug: "mockuser", Icon: "/img/static/img/newuser.png"},
	}
	mockUser = mockUsers[0]

	mockUsersDTO = []dto.User{
		mockUser.ToDto(),
		mockUsers[1].ToDto(),
	}
	mockUserDTO = mockUsersDTO[0]
	userJSON, _ = json.Marshal(mockUserDTO)

	invalidUserDTO = dto.User{UUID: uuid.New(), Name: "", Slug: "mockuser", Icon: "/img/static/img/newuser.png"}
	invalidUserJSON, _ = json.Marshal(invalidUserDTO)
}

func Test_user_service_returns_all_users(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().FindAllUsers().Return(mockUsers, nil)

	result, _ := userService.GetAllUsers()

	if result[0] != mockUsersDTO[0] || result[1] != mockUsersDTO[1] {
		t.Error("Returned users do not match expected")
	}
}

func Test_all_user_service_returns_500_error(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().FindAllUsers().Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := userService.GetAllUsers()

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_update_user_service_returns_updated_user(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().UpdateUser(mockUserDTO).Return(&mockUser, nil)

	result, _ := userService.UpdateUser(userJSON)

	if result.UUID != mockUserDTO.UUID {
		t.Error("Returned users do not match expected")
	}
}

func Test_update_user_service_returns_500_error(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().UpdateUser(mockUserDTO).Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := userService.UpdateUser(userJSON)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_update_user_service_returns_400_error(t *testing.T) {
	setupUserTest(t)
	_, err := userService.UpdateUser(invalidUserJSON)

	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", err.Code)
	}
}

func Test_create_user_service_returns_new_user(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().CreateUser(mockUserDTO).Return(&mockUser, nil)

	result, _ := userService.CreateUser(userJSON)

	if result.UUID != mockUserDTO.UUID {
		t.Error("Returned users do not match expected")
	}
}

func Test_create_user_service_returns_500_error(t *testing.T) {
	setupUserTest(t)

	mockUserRepository.EXPECT().CreateUser(mockUserDTO).Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := userService.CreateUser(userJSON)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_create_user_service_returns_400_error(t *testing.T) {
	setupUserTest(t)
	_, err := userService.UpdateUser(invalidUserJSON)

	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", err.Code)
	}
}

func Test_single_user_service_returns_one_user(t *testing.T) {
	setupUserTest(t)

	mockSlug := "testuser"

	mockUserRepository.EXPECT().FindOneUser(mockSlug).Return(&mockUser, nil)

	result, _ := userService.SingleUser(mockSlug)

	if result.UUID != mockUserDTO.UUID {
		t.Error("Returned users do not match expected")
	}
}

func Test_single_user_service_returns_500_error(t *testing.T) {
	setupUserTest(t)

	mockUser := "testuser"

	mockUserRepository.EXPECT().FindOneUser(mockUser).Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := userService.SingleUser(mockUser)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_delete_user_service_returns_nil(t *testing.T) {
	setupUserTest(t)

	mockSlug := "testuser"

	mockUserRepository.EXPECT().DeleteUser(mockSlug).Return(nil)

	result := userService.DeleteUser(mockSlug)

	if result != nil {
		t.Error("Returned users do not match expected")
	}
}

func Test_delete_user_service_returns_500_error(t *testing.T) {
	setupUserTest(t)

	mockSlug := "testuser"

	mockUserRepository.EXPECT().DeleteUser(mockSlug).Return(errs.NewUnexpectedError("DB error occurred"))

	err := userService.DeleteUser(mockSlug)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}
