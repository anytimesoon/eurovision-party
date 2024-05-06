package service

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"log"
)

//go:generate mockgen -source=userService.go -destination=../../mocks/service/mockUserService.go -package=service eurovision/pkg/service
type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError)
	UpdateUser(dto.User) (*dto.User, *errs.AppError)
	SingleUser(string) (*dto.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dto.User, *errs.AppError)
	Register([]byte) (*dto.NewUser, *errs.AppError)
}

type DefaultUserService struct {
	repo      domain.UserRepository
	broadcast chan dto.SocketMessage
}

func NewUserService(repo domain.UserRepository, broadcast chan dto.SocketMessage) DefaultUserService {
	return DefaultUserService{repo, broadcast}
}

func (service DefaultUserService) GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError) {
	usersDTO := make(map[uuid.UUID]dto.User)

	users, err := service.repo.FindAllUsers()
	if err != nil {
		return usersDTO, err
	}

	for _, user := range users {
		usersDTO[user.UUID] = user.ToDto()
	}

	return usersDTO, nil
}

func (service DefaultUserService) UpdateUser(userDTO dto.User) (*dto.User, *errs.AppError) {
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

func (service DefaultUserService) UpdateUserImage(id uuid.UUID) (*dto.User, *errs.AppError) {

	user, botMessage, appErr := service.repo.UpdateUserImage(id)
	if appErr != nil {
		return nil, appErr
	}

	userDTO := user.ToDto()

	go service.broadcastUserUpdate(userDTO, botMessage)

	return &userDTO, nil
}

func (service DefaultUserService) SingleUser(slug string) (*dto.User, *errs.AppError) {
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

func (service DefaultUserService) GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError) {
	usersDTO := make([]*dto.NewUser, 0)

	users, err := service.repo.FindRegisteredUsers()
	if err != nil {
		return usersDTO, err
	}

	for _, user := range *users {
		usersDTO = append(usersDTO, user.ToDTO())
	}

	return usersDTO, nil
}

func (service DefaultUserService) Register(body []byte) (*dto.NewUser, *errs.AppError) {
	var newUserDTO dto.NewUser
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
func (service DefaultUserService) broadcastNewUser(newUser *dto.NewUser) {
	msg := dto.NewSocketMessage(
		enum.NEW_USER,
		newUser)

	service.broadcast <- msg
}

func (service DefaultUserService) broadcastUserUpdate(user dto.User, comment *dto.Comment) {
	msg := dto.NewSocketMessage(
		enum.UPDATE_USER,
		dto.UpdateMessage{
			UpdatedUser: user,
			Comment:     *comment,
		})

	service.broadcast <- msg
}
