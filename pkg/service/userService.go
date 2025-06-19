package service

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError)
	UpdateUser(dto.User) (*dto.User, *errs.AppError)
	GetOneUser(string) (*dto.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dto.User, *errs.AppError)
	Register(dto.NewUser) (*dto.NewUser, *errs.AppError)
}

type DefaultUserService struct {
	userRepo    data.UserRepository
	authRepo    data.AuthRepository
	broadcast   chan dto.SocketMessage
	commentRepo data.CommentRepository
	voteRepo    data.VoteRepository
}

func NewUserService(
	userRepo data.UserRepository,
	broadcast chan dto.SocketMessage,
	authRepo data.AuthRepositoryDB,
	commentRepo data.CommentRepository,
	voteRepo data.VoteRepository,
) DefaultUserService {
	return DefaultUserService{
		userRepo,
		authRepo,
		broadcast,
		commentRepo,
		voteRepo,
	}
}

func (us DefaultUserService) GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError) {
	usersDTO := make(map[uuid.UUID]dto.User)

	users, err := us.userRepo.GetAllUsers()
	if err != nil {
		return usersDTO, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, user := range users {
		usersDTO[user.UUID] = user.ToDto()
	}

	return usersDTO, nil
}

func (us DefaultUserService) UpdateUser(userDTO dto.User) (*dto.User, *errs.AppError) {
	appErr := userDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	user, err := us.userRepo.GetOneUserById(userDTO.UUID)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	updatedUser, err := us.userRepo.UpdateUser(*dao.User{}.FromDTO(userDTO))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	newUserDTO := updatedUser.ToDto()

	go us.broadcastUserUpdate(newUserDTO, fmt.Sprintf("🤖 %s changed their name to %s", user.Name, updatedUser.Name))

	return &newUserDTO, nil
}

func (us DefaultUserService) UpdateUserImage(id uuid.UUID) (*dto.User, *errs.AppError) {
	user, err := us.userRepo.UpdateUserImage(id)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	userDTO := user.ToDto()

	go us.broadcastUserUpdate(userDTO, fmt.Sprintf("🤖 %s changed their picture", user.Name))

	return &userDTO, nil
}

func (us DefaultUserService) GetOneUser(slug string) (*dto.User, *errs.AppError) {
	user, err := us.userRepo.GetOneUserBySlug(slug)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	userDTO := user.ToDto()

	return &userDTO, nil
}

func (us DefaultUserService) DeleteUser(slug string) *errs.AppError {
	err := us.userRepo.DeleteUser(slug)
	if err != nil {
		return errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return nil
}

func (us DefaultUserService) GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError) {
	usersDTO := make([]*dto.NewUser, 0)

	users, err := us.userRepo.GetRegisteredUsers()
	if err != nil {
		return usersDTO, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, user := range *users {
		auth, err := us.authRepo.GetAuthFromUserId(user.UUID)
		if err != nil {
			log.Println("Failed to get auth from user id.", err)
			continue
		}

		newUser := user.ToNewUserDTO(*auth)

		usersDTO = append(usersDTO, newUser)
	}

	return usersDTO, nil
}

func (us DefaultUserService) Register(newUserDTO dto.NewUser) (*dto.NewUser, *errs.AppError) {
	newUserDTO.Slugify()

	newUser, err := us.userRepo.CreateUser(*dao.User{}.FromNewUserDTO(newUserDTO))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	auth := newUser.GenerateAuth()
	_, err = us.authRepo.CreateAuth(auth)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	err = us.voteRepo.CreateVotes(newUser.UUID)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	createdUserDTO := newUser.ToNewUserDTO(auth)
	go us.broadcastNewUser(createdUserDTO)

	return createdUserDTO, nil
}

func (us DefaultUserService) broadcastNewUser(newUser *dto.NewUser) {
	msg := dto.NewSocketMessage(
		enum.NEW_USER,
		newUser)

	us.broadcast <- msg
}

func (us DefaultUserService) broadcastUserUpdate(user dto.User, comment string) {
	botMessage := dto.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      comment,
		CreatedAt: time.Now(),
	}

	_, err := us.commentRepo.CreateComment(dao.Comment{}.FromDTO(botMessage))
	if err != nil {
		log.Println("Unable to create comment", err)
		return
	}

	msg := dto.NewSocketMessage(
		enum.UPDATE_USER,
		dto.UpdateMessage{
			UpdatedUser: user,
			Comment:     botMessage,
		})

	us.broadcast <- msg
}
