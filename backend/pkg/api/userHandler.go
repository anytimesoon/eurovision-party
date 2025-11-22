package api

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/enum/authLvl"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService  service.UserService
	AssetService service.AssetService
}

func (uh UserHandler) Register(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var newUser *dto.NewUser
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		log.Println("FAILED to parse new user.", err)
		WriteResponse(resp, http.StatusBadRequest, newUser,
			"Failed to parse request body. Please check the format of the request.")
		return
	}

	newUser, appErr = uh.UserService.Register(*newUser)

	if appErr != nil {
		WriteResponse(resp, appErr.Code, newUser, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, newUser, "")
	}
}

func (uh UserHandler) GetAllUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := uh.UserService.GetAllUsers()
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

	if req.Context().Value("auth").(dto.Auth).AuthLvl == authLvl.ADMIN ||
		req.Context().Value("auth").(dto.Auth).UserId == user.UUID {

		user, appErr = uh.UserService.UpdateUser(*user)
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

		user, appErr = uh.UserService.UpdateUserImage(id)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, appErr.Code, user, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
	}
}

func (uh UserHandler) GetOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.UserService.GetOneUser(params["slug"])
	if err != nil {
		WriteResponse(resp, err.Code, user, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
	}
}

func (uh UserHandler) BanUser(resp http.ResponseWriter, req *http.Request) {
	var user *dto.NewUser
	var appErr *errs.AppError

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE.", err)
		return
	}

	user, err = dto.Deserialize[dto.NewUser](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).AuthLvl == authLvl.ADMIN {
		user, appErr = uh.UserService.BanUser(*user)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}
	if appErr != nil {
		WriteResponse(resp, appErr.Code, &dto.User{}, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, user, "")
	}
}

func (uh UserHandler) GetRegisteredUsers(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	users, err := uh.UserService.GetRegisteredUsers(params["userId"])
	if err != nil {
		WriteResponse(resp, err.Code, users, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, users, "")
	}
}
