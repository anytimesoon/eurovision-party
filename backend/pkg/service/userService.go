package service

import (
	"fmt"
	"log"
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	data2 "github.com/anytimesoon/eurovision-party/pkg/data"
	dao2 "github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/enum/chatMsgType"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto2.User, *errs.AppError)
	UpdateUser(dto2.User) (*dto2.User, *errs.AppError)
	GetOneUser(string) (*dto2.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers(string) ([]*dto2.NewUser, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dto2.User, *errs.AppError)
	Register(dto2.NewUser) (*dto2.NewUser, *errs.AppError)
}

type DefaultUserService struct {
	userRepo    data2.UserRepository
	authRepo    data2.AuthRepository
	broadcast   chan dto2.SocketMessage
	commentRepo data2.CommentRepository
	voteRepo    data2.VoteRepository
}

func NewUserService(
	userRepo data2.UserRepository,
	broadcast chan dto2.SocketMessage,
	authRepo data2.AuthRepositoryDB,
	commentRepo data2.CommentRepository,
	voteRepo data2.VoteRepository,
) DefaultUserService {
	return DefaultUserService{
		userRepo,
		authRepo,
		broadcast,
		commentRepo,
		voteRepo,
	}
}

func (us DefaultUserService) GetAllUsers() (map[uuid.UUID]dto2.User, *errs.AppError) {
	usersDTO := make(map[uuid.UUID]dto2.User)

	users, err := us.userRepo.GetAllUsers()
	if err != nil {
		return usersDTO, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, user := range users {
		usersDTO[user.UUID] = user.ToDto()
	}

	return usersDTO, nil
}

func (us DefaultUserService) UpdateUser(userDTO dto2.User) (*dto2.User, *errs.AppError) {
	appErr := userDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	user, err := us.userRepo.GetOneUserById(userDTO.UUID)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	updatedUser, err := us.userRepo.UpdateUser(*dao2.User{}.FromDTO(userDTO))
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	newUserDTO := updatedUser.ToDto()

	go us.broadcastUserUpdate(newUserDTO, fmt.Sprintf("🤖 %s changed their name to %s", user.Name, updatedUser.Name))

	return &newUserDTO, nil
}

func (us DefaultUserService) UpdateUserImage(id uuid.UUID) (*dto2.User, *errs.AppError) {
	user, err := us.userRepo.UpdateUserImage(id)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	userDTO := user.ToDto()

	go us.broadcastUserUpdate(userDTO, fmt.Sprintf("🤖 %s changed their picture", user.Name))

	return &userDTO, nil
}

func (us DefaultUserService) GetOneUser(slug string) (*dto2.User, *errs.AppError) {
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

func (us DefaultUserService) GetRegisteredUsers(userId string) ([]*dto2.NewUser, *errs.AppError) {
	usersDTO := make([]*dto2.NewUser, 0)

	requestingUserId, err := uuid.Parse(userId)
	if err != nil {
		return usersDTO, errs.NewUnexpectedError("Could not parse user id.")
	}

	requestingUser, err := us.userRepo.GetOneUserById(requestingUserId)
	if err != nil {
		return usersDTO, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	users, err := us.userRepo.GetRegisteredUsers(*requestingUser)
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

func (us DefaultUserService) Register(newUserDTO dto2.NewUser) (*dto2.NewUser, *errs.AppError) {
	requestingUser, err := us.userRepo.GetOneUserById(newUserDTO.CreatedBy)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	if !requestingUser.CanInvite {
		return nil, errs.NewForbiddenError(errs.Common.MaxInvitesExceeded)
	}

	newUserDTO.Slugify()

	newUser, err := us.userRepo.CreateUser(*dao2.User{}.FromNewUserDTO(newUserDTO, requestingUser))
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

	requestingUser.Invites = append(requestingUser.Invites, newUser.UUID)
	if requestingUser.AuthLvl != authLvl.ADMIN && uint8(len(requestingUser.Invites)) >= conf.App.MaxInvites {
		requestingUser.CanInvite = false
	}
	_, err = us.userRepo.UpdateUser(*requestingUser)
	if err != nil {
		log.Println("Failed to update requesting user invites.", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return createdUserDTO, nil
}

func (us DefaultUserService) broadcastNewUser(newUser *dto2.NewUser) {
	msg := dto2.NewSocketMessage(
		chatMsgType.NEW_USER,
		newUser)

	us.broadcast <- msg
}

func (us DefaultUserService) broadcastUserUpdate(user dto2.User, comment string) {
	botMessage := dto2.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      comment,
		CreatedAt: time.Now(),
	}

	_, err := us.commentRepo.CreateComment(dao2.Comment{}.FromDTO(botMessage))
	if err != nil {
		log.Println("Unable to create comment", err)
		return
	}

	msg := dto2.NewSocketMessage(
		chatMsgType.UPDATE_USER,
		dto2.UpdateMessage{
			UpdatedUser: user,
			Comment:     botMessage,
		})

	us.broadcast <- msg
}
