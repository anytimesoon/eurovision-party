package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"log"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
	UpdateUser([]byte) (dto.User, error)
	CreateUser([]byte) (dto.User, error)
	SingleUser(string) (dto.User, error)
	DeleteUser(string) error
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}

func (service DefaultUserService) GetAllUsers() ([]dto.User, error) {
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

func (service DefaultUserService) UpdateUser(body []byte) (dto.User, error) {
	var userDTO dto.User
	err := json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return userDTO, err
	}

	user, err := service.repo.UpdateUser(userDTO)
	if err != nil {
		log.Println("FAILED to update user", err)
		return userDTO, err
	}

	return user.ToDto(), nil
}

func (service DefaultUserService) CreateUser(body []byte) (dto.User, error) {
	var userDTO dto.User
	err := json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return userDTO, err
	}

	user, err := service.repo.CreateUser(userDTO)
	if err != nil {
		log.Println("FAILED to create user", err)
		return userDTO, err
	}

	return user.ToDto(), nil
}

func (service DefaultUserService) SingleUser(slug string) (dto.User, error) {
	user, err := service.repo.FindOneUser(slug)
	if err != nil {
		log.Println("FAILED to find user", err)
	}

	return user.ToDto(), nil
}

func (service DefaultUserService) DeleteUser(slug string) error {
	err := service.repo.DeleteUser(slug)
	if err != nil {
		log.Println("FAILED to delete user", err)
	}

	return nil
}
