package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"io"
	"log"
	"time"
)

type AuthService interface {
	Login(dto.Auth) (*dto.Auth, *dto.User, *errs.AppError)
	Authorize(string) (*dto.Auth, *errs.AppError)
}

type DefaultAuthService struct {
	repo data.AuthRepositoryDB
}

var secretKey []byte

func init() {
	var err error
	// TODO: create random string
	secretKey, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}
}

func NewAuthService(repo data.AuthRepositoryDB) DefaultAuthService {
	return DefaultAuthService{repo}
}

func (das DefaultAuthService) Login(authDTO dto.Auth) (*dto.Auth, *dto.User, *errs.AppError) {
	auth, user, appErr := das.repo.Login(&authDTO)
	if appErr != nil {
		return nil, nil, appErr
	}

	authJson, err := json.Marshal(auth.ToDTO())
	if err != nil {
		log.Printf("Failed to marshall auth %+v %s", auth, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	e, err := encrypt(string(authJson))
	if err != nil {
		log.Printf("Couldn't encrypt the creds for %+v", auth)
		return nil, nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	returnableAuth := auth.ToReturnableDTO(e)
	userDTO := user.ToDto()

	return &returnableAuth, &userDTO, nil
}

func (das DefaultAuthService) Authorize(token string) (*dto.Auth, *errs.AppError) {
	authDTO, appErr := decrypt(token)
	if appErr != nil {
		return nil, appErr
	}
	log.Printf("Session %+v", authDTO)
	if authDTO.Expiration.Before(time.Now()) {
		log.Printf("Session has expired")
		return nil, errs.NewUnauthorizedError(errs.Common.Login)
	}

	_, appErr = das.repo.Authorize(authDTO)
	if appErr != nil {
		return nil, appErr
	}

	return authDTO, nil
}

func encrypt(auth string) (string, error) {
	// Create a new AES cipher block from the secret key.
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	// Wrap the cipher block in Galois Counter Mode.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a unique nonce containing 12 random bytes.
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	plaintext := fmt.Sprintf("%+v", auth)

	// Encrypt the data using aesGCM.Seal(). By passing the nonce as the first
	// parameter, the encrypted data will be appended to the nonce — meaning
	// that the returned encryptedValue variable will be in the format
	// "{nonce}{encrypted plaintext data}".
	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	base64Value := base64.RawURLEncoding.EncodeToString(encryptedValue)
	return base64Value, nil
}

func decrypt(base64Token string) (*dto.Auth, *errs.AppError) {
	token, err := base64.RawURLEncoding.DecodeString(base64Token)
	if err != nil {
		log.Println("Failed to decode base 64", err)
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	// Create a new AES cipher block from the secret key.
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Println("Failed to create new cipher", err)
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	// Wrap the cipher block in Galois Counter Mode.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Failed to wrap cipher in block", err)
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	// Get the nonce size.
	nonceSize := aesGCM.NonceSize()

	// To avoid a potential 'index out of range' panic in the next step, we
	// check that the length of the encrypted value is at least the nonce
	// size.
	if len(token) < nonceSize {
		log.Println("Nonce was too large")
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	// Split apart the nonce from the actual encrypted data.
	nonce := token[:nonceSize]
	ciphertext := token[nonceSize:]

	// Use aesGCM.Open() to decrypt and authenticate the data. If this fails,
	// return a ErrInvalidValue error.
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("Failed to decrypt token %s", err)
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	var auth dto.Auth
	err = json.Unmarshal(plaintext, &auth)
	if err != nil {
		log.Printf("Failed to unmarshal %s token %s", plaintext, err)
		return nil, errs.NewUnexpectedError(errs.Common.Login)
	}

	// Return the plaintext cookie value.
	return &auth, nil
}
