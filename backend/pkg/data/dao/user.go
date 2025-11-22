package dao

import (
	"time"

	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	AuthLvl   authLvl.AuthLvl
	Name      string
	Slug      string `boltholdUnique:"UniqueSlug"`
	Icon      string
	Invites   []uuid.UUID
	CreatedBy uuid.UUID `boltholdIndex:"CreatedBy"`
	CanInvite bool
	IsBanned  bool
}

func (u User) ToDto() dto2.User {
	return dto2.User{
		Name:      u.Name,
		Slug:      u.Slug,
		UUID:      u.UUID,
		Icon:      u.Icon,
		AuthLvl:   u.AuthLvl,
		Invites:   u.Invites,
		CreatedBy: u.CreatedBy,
		CanInvite: u.CanInvite,
	}
}

func (u User) FromDTO(userDTO dto2.User) *User {
	return &User{
		UUID:      userDTO.UUID,
		AuthLvl:   userDTO.AuthLvl,
		Name:      userDTO.Name,
		Slug:      userDTO.Slug,
		Icon:      userDTO.Icon,
		Invites:   userDTO.Invites,
		CreatedBy: userDTO.CreatedBy,
		CanInvite: userDTO.CanInvite,
	}
}

func (u User) ToNewUserDTO(auth Auth) *dto2.NewUser {
	return &dto2.NewUser{
		UUID:      u.UUID,
		AuthLvl:   u.AuthLvl,
		Name:      u.Name,
		Slug:      u.Slug,
		Token:     auth.AuthToken,
		CreatedBy: u.CreatedBy,
	}
}

func (u User) FromNewUserDTO(newUser dto2.NewUser, requestingUser *User) *User {
	newAuthLvl := authLvl.USER
	if requestingUser.AuthLvl == authLvl.USER {
		newAuthLvl = authLvl.FRIEND_OF_FRIEND
	}
	return &User{
		UUID:      uuid.New(),
		AuthLvl:   newAuthLvl,
		Name:      newUser.Name,
		Slug:      newUser.Slug,
		Icon:      "default",
		Invites:   make([]uuid.UUID, 0),
		CreatedBy: requestingUser.UUID,
		CanInvite: requestingUser.AuthLvl == authLvl.ADMIN,
	}
}

func (u User) GenerateAuth() Auth {
	a := Auth{
		UserId:       u.UUID,
		AuthTokenExp: time.Now().Add(7 * 24 * time.Hour),
		AuthLvl:      u.AuthLvl,
		Slug:         u.Slug,
	}
	a.GenerateSecureToken(40)
	return a
}
