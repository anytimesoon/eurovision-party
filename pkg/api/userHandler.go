package api

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
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

func (uh UserHandler) Register(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var auth *dto.NewUser
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.ADMIN {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		auth, appErr = uh.Service.Register(body)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, appErr.Code, auth, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, auth, "")
	}
}

func (uh UserHandler) FindAllUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uh.Service.GetAllUsers()
	if err != nil {
		WriteResponse(resp, err.Code, users, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, users, "")
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

	user, err = dto.Deserialize[dto.User](body)
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
		WriteResponse(resp, appErr.Code, user, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
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
			WriteResponse(resp, appErr.Code, user, appErr.Message)
		}

		user, appErr = uh.Service.UpdateUserImage(id)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, appErr.Code, user, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
	}
}

func (uh UserHandler) FindOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.Service.SingleUser(params["slug"])
	if err != nil {
		WriteResponse(resp, err.Code, user, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
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
		WriteResponse(resp, err.Code, &dto.User{}, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, &dto.User{}, "")
	}
}

func (uh UserHandler) FindRegisteredUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uh.Service.GetRegisteredUsers()
	if err != nil {
		WriteResponse(resp, err.Code, users, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, users, "")
	}
}
