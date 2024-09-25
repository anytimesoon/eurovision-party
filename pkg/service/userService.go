package service

import (
	"encoding/json"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"log"
)

type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto2.User, *errs.AppError)
	UpdateUser(dto2.User) (*dto2.User, *errs.AppError)
	SingleUser(string) (*dto2.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers() ([]*dto2.NewUser, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dto2.User, *errs.AppError)
	Register([]byte) (*dto2.NewUser, *errs.AppError)
}

type DefaultUserService struct {
	repo      data.UserRepository
	broadcast chan dto2.SocketMessage
}

func NewUserService(repo data.UserRepository, broadcast chan dto2.SocketMessage) DefaultUserService {
	return DefaultUserService{repo, broadcast}
}

func (service DefaultUserService) GetAllUsers() (map[uuid.UUID]dto2.User, *errs.AppError) {
	usersDTO := make(map[uuid.UUID]dto2.User)

	users, err := service.repo.FindAllUsers()
	if err != nil {
		return usersDTO, err
	}

	for _, user := range users {
		usersDTO[user.UUID] = user.ToDto()
	}

	return usersDTO, nil
}

func (service DefaultUserService) UpdateUser(userDTO dto2.User) (*dto2.User, *errs.AppError) {
	appErr := userDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	user, botMessage, appErr := service.repo.UpdateUser(userDTO)
	if appErr != nil {
		return nil, appErr
	}

	newUserDTO := user.ToDto()

	go service.broadcastUserUpdate(newUserDTO, botMessage)

	return &newUserDTO, nil
}

func (service DefaultUserService) UpdateUserImage(id uuid.UUID) (*dto2.User, *errs.AppError) {

	user, botMessage, appErr := service.repo.UpdateUserImage(id)
	if appErr != nil {
		return nil, appErr
	}

	userDTO := user.ToDto()

	go service.broadcastUserUpdate(userDTO, botMessage)

	return &userDTO, nil
}

func (service DefaultUserService) SingleUser(slug string) (*dto2.User, *errs.AppError) {
	user, err := service.repo.FindOneUser(slug)
	if err != nil {
		return nil, err
	}

	userDTO := user.ToDto()

	return &userDTO, nil
}

func (service DefaultUserService) DeleteUser(slug string) *errs.AppError {
	err := service.repo.DeleteUser(slug)
	if err != nil {
		return err
	}

	return nil
}

func (service DefaultUserService) GetRegisteredUsers() ([]*dto2.NewUser, *errs.AppError) {
	usersDTO := make([]*dto2.NewUser, 0)

	users, err := service.repo.FindRegisteredUsers()
	if err != nil {
		return usersDTO, err
	}

	for _, user := range *users {
		usersDTO = append(usersDTO, user.ToDTO())
	}

	return usersDTO, nil
}

func (service DefaultUserService) Register(body []byte) (*dto2.NewUser, *errs.AppError) {
	var newUserDTO dto2.NewUser
	err := json.Unmarshal(body, &newUserDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	newUserDTO.Slugify()

	// create new user
	newUser, appErr := service.repo.CreateUser(newUserDTO)
	if appErr != nil {
		log.Println("Failed to create user", appErr)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	createdUserDTO := newUser.ToDTO()
	go service.broadcastNewUser(createdUserDTO)

	return createdUserDTO, nil
}
func (service DefaultUserService) broadcastNewUser(newUser *dto2.NewUser) {
	msg := dto2.NewSocketMessage(
		enum.NEW_USER,
		newUser)

	service.broadcast <- msg
}

func (service DefaultUserService) broadcastUserUpdate(user dto2.User, comment *dto2.Comment) {
	msg := dto2.NewSocketMessage(
		enum.UPDATE_USER,
		dto2.UpdateMessage{
			UpdatedUser: user,
			Comment:     *comment,
		})

	service.broadcast <- msg
}
