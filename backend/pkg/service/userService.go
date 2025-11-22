package service

import (
	"fmt"
	"log"
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/enum/chatMsgType"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError)
	UpdateUser(dto.User) (*dto.User, *errs.AppError)
	GetOneUser(string) (*dto.User, *errs.AppError)
	BanUser(dto.NewUser) (*dto.NewUser, *errs.AppError)
	GetRegisteredUsers(string) ([]*dto.NewUser, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dto.User, *errs.AppError)
	Register(dto.NewUser) (*dto.NewUser, *errs.AppError)
}

type DefaultUserService struct {
	userRepo    data.UserRepository
	authRepo    data.AuthRepository
	sessionRepo data.SessionRepository
	broadcast   chan dto.SocketMessage
	commentRepo data.CommentRepository
	voteRepo    data.VoteRepository
}

func NewUserService(
	userRepo data.UserRepository,
	broadcast chan dto.SocketMessage,
	authRepo data.AuthRepositoryDB,
	sessionRepo data.SessionRepositoryDB,
	commentRepo data.CommentRepository,
	voteRepo data.VoteRepository,
) DefaultUserService {
	return DefaultUserService{
		userRepo,
		authRepo,
		sessionRepo,
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

func (us DefaultUserService) BanUser(newUser dto.NewUser) (*dto.NewUser, *errs.AppError) {
	user, err := us.userRepo.GetOneUserBySlug(newUser.Slug)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	auth, err := us.authRepo.GetAuthFromUserId(user.UUID)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	session, err := us.sessionRepo.GetSession(auth.SessionToken)
	if err != nil {
		session = &dao.Session{}
	}

	if auth.SessionToken != "" {
		err = us.sessionRepo.DeleteSession(auth.SessionToken)
		if err != nil {
			return nil, errs.NewUnexpectedError(errs.Common.DBFail)
		}
		log.Println("Deleted session for user", user.Name)
	}

	err = us.authRepo.DeleteAuth(auth.AuthToken)
	if err != nil {
		log.Println("Failed to delete auth for user", user.Name)
		_ = us.sessionRepo.UpsertSession(session.AuthToken, session.SessionToken, user.UUID)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}
	log.Println("Deleted auth for user", user.Name)

	user.IsBanned = true
	_, err = us.userRepo.UpdateUser(*user)
	if err != nil {
		log.Println("Failed to ban user", err)
		_ = us.sessionRepo.UpsertSession(session.AuthToken, session.SessionToken, user.UUID)
		_, _ = us.authRepo.CreateAuth(*auth)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	us.broadcastUserUpdate(user.ToDto(), fmt.Sprintf("🤖 %s has been banned 🚫", user.Name))

	return &newUser, nil
}

func (us DefaultUserService) GetRegisteredUsers(userId string) ([]*dto.NewUser, *errs.AppError) {
	usersDTO := make([]*dto.NewUser, 0)

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

func (us DefaultUserService) Register(newUserDTO dto.NewUser) (*dto.NewUser, *errs.AppError) {
	requestingUser, err := us.userRepo.GetOneUserById(newUserDTO.CreatedBy)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	if !requestingUser.CanInvite {
		return nil, errs.NewForbiddenError(errs.Common.MaxInvitesExceeded)
	}

	newUserDTO.Slugify()

	newUser, err := us.userRepo.CreateUser(*dao.User{}.FromNewUserDTO(newUserDTO, requestingUser))
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

func (us DefaultUserService) broadcastNewUser(newUser *dto.NewUser) {
	msg := dto.NewSocketMessage(
		chatMsgType.NEW_USER,
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
		chatMsgType.UPDATE_USER,
		dto.UpdateMessage{
			UpdatedUser: user,
			Comment:     botMessage,
		})

	us.broadcast <- msg
}
