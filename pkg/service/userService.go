package service

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"github.com/google/uuid"
)

//go:generate mockgen -source=userService.go -destination=../../mocks/service/mockUserService.go -package=service eurovision/pkg/service
type UserService interface {
	GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError)
	UpdateUser(dto.User) (*dto.User, *errs.AppError)
	SingleUser(string) (*dto.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	GetRegisteredUsers() ([]*dto.NewUser, *errs.AppError)
	UpdateUserImage(image dto.UserImage) (*dto.User, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}

func (service DefaultUserService) GetAllUsers() (map[uuid.UUID]dto.User, *errs.AppError) {
	usersDTO := make(map[uuid.UUID]dto.User, 0)

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

	user, appErr := service.repo.UpdateUser(userDTO)
	if appErr != nil {
		return nil, appErr
	}

	userDTO = user.ToDto()

	return &userDTO, nil
}

func (service DefaultUserService) UpdateUserImage(imageDTO dto.UserImage) (*dto.User, *errs.AppError) {
	image, appErr := stringToBin(imageDTO.Image)
	if appErr != nil {
		return nil, appErr
	}

	fileLocation, appErr := cropImage(imageDTO, image)
	if appErr != nil {
		return nil, appErr
	}

	imageDTO.Image = fileLocation

	user, appErr := service.repo.UpdateUserImage(imageDTO)
	if appErr != nil {
		return nil, appErr
	}

	userDTO := user.ToDto()

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
