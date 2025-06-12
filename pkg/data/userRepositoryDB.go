package data

import (
	"errors"
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
	"log"
	"strconv"
	"time"
)

type UserRepository interface {
	GetUser(string) (*dao.User, error)

	// CreateUser OLD
	CreateUser(dto.NewUser) (*dao.NewUser, *errs.AppError)
	FindAllUsers() ([]dao.User, *errs.AppError)
	FindOneUser(string) (*dao.User, *errs.AppError)
	DeleteUser(string) *errs.AppError
	FindRegisteredUsers() (*[]dao.NewUser, *errs.AppError)
	UpdateUser(dto.User) (*dao.User, *dto.Comment, *errs.AppError)
	UpdateUserImage(uuid.UUID) (*dao.User, *dto.Comment, *errs.AppError)
	VerifySlug(*dto.NewUser) error
}

type UserRepositoryDb struct {
	store *bolthold.Store
}

func NewUserRepositoryDb(store *bolthold.Store) UserRepositoryDb {
	return UserRepositoryDb{store}
}

func (db UserRepositoryDb) GetUser(userId string) (*dao.User, error) {
	var user dao.User
	err := db.store.Get(userId, &user)
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, err
	}
	return &user, nil
}

func (db UserRepositoryDb) CreateUser(userDTO dto.NewUser) (*dao.NewUser, *errs.AppError) {
	var newUser dao.NewUser

	err := db.VerifySlug(&userDTO)
	if err != nil {
		log.Printf("Error when slufigying user %s with message %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	newUser.FromDTO(userDTO)
	auth := newUser.GenerateAuth()
	user := newUser.ToUser()
	err = db.store.Insert(user.UUID.String(), user)
	if err != nil {
		log.Printf("Error when creating new user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "user")
	}

	err = db.store.Insert(auth.AuthToken, auth)
	if err != nil {
		log.Printf("Error when creating new auth for user %s, %s", userDTO.Name, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "a new user")
	}

	return &newUser, nil
}

func (db UserRepositoryDb) VerifySlug(userDTO *dto.NewUser) error {
	// Verify the name is unique or add a number to the end
	counter := 0
	for {
		var user dao.User
		if counter > 0 {
			userDTO.Slug = userDTO.Slug + "-" + strconv.Itoa(counter)
		}

		err := db.store.FindOne(&user, bolthold.Where("Slug").Eq(userDTO.Slug))
		if err != nil {
			if errors.Is(err, bolthold.ErrNotFound) {
				// no users with this slug found, so validation is complete
				return nil
			}
			return err
		}

		counter++
	}
}

func (db UserRepositoryDb) FindAllUsers() ([]dao.User, *errs.AppError) {
	users := make([]dao.User, 0)

	//query := "SELECT * FROM user"
	query := &bolthold.Query{}
	err := db.store.Find(&users, query)
	if err != nil {
		log.Println("Error while querying user table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return users, nil
}

func (db UserRepositoryDb) UpdateUser(userDTO dto.User) (*dao.User, *dto.Comment, *errs.AppError) {
	var user dao.User

	err := db.store.Get(userDTO.UUID.String(), &user)
	if err != nil {
		log.Printf("Error while fetching user %s after update %s", userDTO.Name, err)
		return nil, nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	user.Name = userDTO.Name
	err = db.store.Update(user.UUID.String(), user)
	if err != nil {
		log.Println("Error while updating user table", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "user")
	}

	botComment := dao.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      fmt.Sprintf("🤖 %s changed their name to %s", user.Name, userDTO.Name),
		CreatedAt: time.Now(),
	}
	err = db.store.Insert(botComment.UUID.String(), botComment)
	if err != nil {
		log.Printf("Error while writing bot comment after name update %s", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	botCommentDTO := botComment.ToDto()
	return &user, &botCommentDTO, nil
}

func (db UserRepositoryDb) UpdateUserImage(id uuid.UUID) (*dao.User, *dto.Comment, *errs.AppError) {
	var user dao.User

	err := db.store.Get(id.String(), &user)
	if err != nil {
		log.Printf("Error while fetching user %s to image %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	user.Icon = id.String() + ".png"
	//_, err = tx.Exec(updateUserImageQuery, id.String()+".png", id)
	err = db.store.Update(id.String(), user)
	if err != nil {
		log.Printf("Error while updating user image for user %s. %s", id, err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	var botComment = dto.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      fmt.Sprintf("🤖 %s changed their picture", user.Name),
		CreatedAt: time.Now(),
	}

	err = db.store.Insert(botComment.UUID.String(), botComment)
	if err != nil {
		log.Printf("Error while writing bot comment after updating image %s", err)
		return nil, nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "image")
	}

	return &user, &botComment, nil
}

func (db UserRepositoryDb) FindOneUser(slug string) (*dao.User, *errs.AppError) {
	var user dao.User

	err := db.store.FindOne(&user, bolthold.Where("slug").Eq(slug))
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	return &user, nil
}

func (db UserRepositoryDb) DeleteUser(slug string) *errs.AppError {
	var user dao.User

	err := db.store.FindOne(&user, bolthold.Where("slug").Eq(slug))
	if err != nil {
		log.Printf("Error when fetching user: %s", err)
		return errs.NewNotFoundError(errs.Common.NotFound + "user")
	}

	err = db.store.Delete(user.UUID.String(), user)
	if err != nil {
		log.Println("Error when deleting user", err)
		return errs.NewUnexpectedError(errs.Common.NotDeleted + "user")
	}

	return nil
}

func (db UserRepositoryDb) FindRegisteredUsers() (*[]dao.NewUser, *errs.AppError) {
	users := make([]dao.User, 0)
	newUsers := make([]dao.NewUser, 0)

	err := db.store.Find(&users, bolthold.Where("AuthLvl").Ne(enum.BOT))
	if err != nil {
		log.Println("Error while querying user table for registered users", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, user := range users {
		var auth dao.Auth
		newUser := user.ToNewUser()

		err = db.store.FindOne(&auth, bolthold.Where("UserId").Eq(user.UUID))
		if err != nil {
			log.Println("Error while querying auth table for registered users", err)
			return nil, errs.NewUnexpectedError(errs.Common.DBFail)
		}

		newUser.Token = auth.AuthToken
		newUsers = append(newUsers, *newUser)
	}

	return &newUsers, nil
}
