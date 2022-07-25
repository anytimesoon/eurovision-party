package service

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
	// UpdateUser([]byte) (dto.User, error)
	// SingleUser(string) (dto.User, error)
	// DeleteUser(string)
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
