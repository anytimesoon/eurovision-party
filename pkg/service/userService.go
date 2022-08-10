package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

type UserService interface {
	GetAllUsers() ([]dto.User, *errs.AppError)
	UpdateUser([]byte) (*dto.User, *errs.AppError)
	CreateUser([]byte) (*dto.User, *errs.AppError)
	SingleUser(string) (*dto.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}

func (service DefaultUserService) GetAllUsers() ([]dto.User, *errs.AppError) {
	usersDTO := make([]dto.User, 0)

	users, err := service.repo.FindAllUsers()
	if err != nil {
		return usersDTO, err
	}

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToDto())
	}

	return usersDTO, nil
}

func (service DefaultUserService) UpdateUser(body []byte) (*dto.User, *errs.AppError) {
	var userDTO dto.User
	err := json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

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

func (service DefaultUserService) CreateUser(body []byte) (*dto.User, *errs.AppError) {
	var userDTO dto.User
	err := json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	userDTO.Slugify()

	appErr := userDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	user, appErr := service.repo.CreateUser(userDTO)
	if appErr != nil {
		return nil, appErr
	}

	userDTO = user.ToDto()

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
