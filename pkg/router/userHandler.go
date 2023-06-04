package router

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"eurovision/pkg/errs"
	"eurovision/pkg/service"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service service.UserService
}

func (uh UserHandler) FindAllUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uh.Service.GetAllUsers()
	if err != nil {
		writeResponse(resp, req, err.Code, users, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, users, "")
	}
}

func (uh UserHandler) UpdateUser(resp http.ResponseWriter, req *http.Request) {
	var user *dto.User
	var appErr *errs.AppError

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!", err)
		return
	}

	user, err = dto.Decode[dto.User](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin ||
		req.Context().Value("auth").(dto.Auth).UserId == user.UUID {

		user, appErr = uh.Service.UpdateUser(*user)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, *user, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *user, "")
	}
}

func (uh UserHandler) UpdateImage(resp http.ResponseWriter, req *http.Request) {
	var userImage *dto.UserImage
	var user *dto.User
	var appErr *errs.AppError

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER IMAGE UPDATE!", err)
		return
	}

	userImage, err = dto.Decode[dto.UserImage](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin ||
		req.Context().Value("auth").(dto.Auth).UserId == userImage.UUID {
		user, appErr = uh.Service.UpdateUserImage(*userImage)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, *user, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *user, "")
	}
}

func (uh UserHandler) FindOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.Service.SingleUser(params["slug"])
	if err != nil {
		writeResponse(resp, req, err.Code, *user, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *user, "")
	}
}

func (uh UserHandler) RemoveUser(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin {
		params := mux.Vars(req)
		err = uh.Service.DeleteUser(params["slug"])
	} else {
		err = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}
	if err != nil {
		writeResponse(resp, req, err.Code, dto.User{}, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, dto.User{}, "")
	}
}

func (uh UserHandler) FindRegisteredUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uh.Service.GetRegisteredUsers()
	if err != nil {
		writeResponse(resp, req, err.Code, users, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, users, "")
	}
}
