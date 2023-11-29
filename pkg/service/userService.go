package service

import (
	"encoding/json"
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
	"log"
	"time"
)

//go:generate mockgen -source=userService.go -destination=../../mocks/service/mockUserService.go -package=service eurovision/pkg/service
type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError)
	UpdateUser(dto.User) (*dto.User, *errs.AppError)
	SingleUser(string) (*dto.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError)
	UpdateUserImage(image dto.UserAvatar) (*dto.User, *errs.AppError)
}

type DefaultUserService struct {
	repo      domain.UserRepository
	broadcast chan []byte
}

func NewUserService(repo domain.UserRepository, broadcast chan []byte) DefaultUserService {
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

	oldUser, updatedUser, appErr := service.repo.UpdateUser(userDTO)
	if appErr != nil {
		return nil, appErr
	}

	newUserDTO := updatedUser.ToDto()

	go service.broadcastUserUpdate(newUserDTO, fmt.Sprintf("🤖 %s changed their name to %s", oldUser.Name, updatedUser.Name))

	return &newUserDTO, nil
}

func (service DefaultUserService) UpdateUserImage(avatarDTO dto.UserAvatar) (*dto.User, *errs.AppError) {
	img, appErr := cropImage(&avatarDTO)
	if appErr != nil {
		return nil, appErr
	}

	appErr = storeImageToDisk(img)
	if appErr != nil {
		return nil, appErr
	}

	user, appErr := service.repo.UpdateUserImage(avatarDTO, img)
	if appErr != nil {
		return nil, appErr
	}

	userDTO := user.ToDto()

	go service.broadcastUserUpdate(userDTO, fmt.Sprintf("🤖 %s changed their picture", userDTO.Name))

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

func (service DefaultUserService) broadcastUserUpdate(newUser dto.User, commentText string) {
	updateMessage := dto.UpdateMessage{
		UpdatedUser: newUser,
		Comment: dto.Comment{
			UUID:      uuid.New(),
			UserId:    conf.App.BotId,
			Text:      commentText,
			CreatedAt: time.Now(),
		},
	}

	msg, err := json.Marshal(updateMessage)
	if err != nil {
		log.Println("Failed to send update updateMessage")
	}

	service.broadcast <- msg
}
