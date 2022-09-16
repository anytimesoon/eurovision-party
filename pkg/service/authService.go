package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

type AuthService interface {
	Login() *errs.AppError
	Token([]byte) ([]byte, *errs.AppError)
	Register([]byte) (*dto.Auth, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepositoryDB
}

func NewAuthService(repo domain.AuthRepositoryDB) DefaultAuthService {
	return DefaultAuthService{repo}
}

func (das DefaultAuthService) Login() *errs.AppError {
	return errs.NewInvalidError("asdf")
}

func (das DefaultAuthService) Token(body []byte) ([]byte, *errs.AppError) {
	return []byte("hello"), nil
}

func (das DefaultAuthService) Register(body []byte) (*dto.Auth, *errs.AppError) {
	var newUserDTO dto.NewUser
	err := json.Unmarshal(body, &newUserDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	// verify user doesn't already exist
	user := das.repo.FindOneUserByEmail(newUserDTO.Email)
	if user.Email == newUserDTO.Email {
		log.Printf("User with email %s alread exists", newUserDTO.Email)
		return nil, errs.NewUnexpectedError("User with email " + newUserDTO.Email + " alread exists")
	}

	newUserDTO.Slugify()

	// create new user
	auth, appErr := das.repo.CreateUser(newUserDTO)
	if appErr != nil {
		log.Println("Failed to create user", appErr)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	authDTO := auth.ToDTO()

	return &authDTO, nil
}
