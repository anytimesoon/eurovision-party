package router

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type UserHandler struct {
	Service      service.UserService
	AssetService service.AssetService
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
		log.Println("FAILED to read body of USER UPDATE.", err)
		return
	}

	user, err = dto.Decode[dto.User](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.ADMIN ||
		req.Context().Value("auth").(dto.Auth).UserId == user.UUID {

		user, appErr = uh.Service.UpdateUser(*user)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, user, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, user, "")
	}
}

func (uh UserHandler) UpdateImage(resp http.ResponseWriter, req *http.Request) {
	var user *dto.User
	var appErr *errs.AppError

	log.Println("Starting image save")

	err := req.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		log.Println("Failed to parse form data", err)
		return
	}

	defer func(MultipartForm *multipart.Form) {
		err := MultipartForm.RemoveAll()
		if err != nil {

		}
	}(req.MultipartForm)

	id, err := uuid.Parse(req.PostFormValue("id"))
	if err != nil {
		log.Println("Failed to parse user id", err)
	}

	if req.Context().Value("auth").(dto.Auth).UserId == id {

		fileHeaders := req.MultipartForm.File["file"]
		appErr = uh.AssetService.PersistImage(fileHeaders, filepath.Join(conf.App.Assets, "avatars"))
		if appErr != nil {
			writeResponse(resp, req, appErr.Code, user, appErr.Message)
		}

		user, appErr = uh.Service.UpdateUserImage(id)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, user, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, user, "")
	}
}

func (uh UserHandler) FindOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.Service.SingleUser(params["slug"])
	if err != nil {
		writeResponse(resp, req, err.Code, user, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, user, "")
	}
}

func (uh UserHandler) RemoveUser(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.ADMIN {
		params := mux.Vars(req)
		err = uh.Service.DeleteUser(params["slug"])
	} else {
		err = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}
	if err != nil {
		writeResponse(resp, req, err.Code, &dto.User{}, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, &dto.User{}, "")
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
