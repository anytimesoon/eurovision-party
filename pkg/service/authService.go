package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"io"
	"log"
)

type AuthService interface {
	Login([]byte) (*dto.EAuth, *errs.AppError)
	Token([]byte) ([]byte, *errs.AppError)
	Register([]byte) (*dto.Auth, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepositoryDB
}

var secretKey []byte

func init() {
	var err error
	secretKey, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}
}

func NewAuthService(repo domain.AuthRepositoryDB) DefaultAuthService {
	return DefaultAuthService{repo}
}

func (das DefaultAuthService) Login(body []byte) (*dto.EAuth, *errs.AppError) {
	var authDTO dto.Auth
	err := json.Unmarshal(body, &authDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	auth, appErr := das.repo.Authenticate(&authDTO)
	if appErr != nil {
		return nil, appErr
	}

	e, err := encrypt(auth.ToDTO())

	eAuth := &dto.EAuth{
		EToken: e,
	}

	return eAuth, nil
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

func encrypt(auth dto.Auth) ([]byte, error) {
	// Create a new AES cipher block from the secret key.
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	// Wrap the cipher block in Galois Counter Mode.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a unique nonce containing 12 random bytes.
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	plaintext := fmt.Sprintf("%+v", auth)

	// Encrypt the data using aesGCM.Seal(). By passing the nonce as the first
	// parameter, the encrypted data will be appended to the nonce — meaning
	// that the returned encryptedValue variable will be in the format
	// "{nonce}{encrypted plaintext data}".
	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	base64Value := base64.URLEncoding.EncodeToString(encryptedValue)
	return []byte(base64Value), nil
}
