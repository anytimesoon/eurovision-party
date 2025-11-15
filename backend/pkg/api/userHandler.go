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
	service2 "github.com/anytimesoon/eurovision-party/pkg/service"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService  service2.UserService
	AssetService service2.AssetService
}

func (uh UserHandler) Register(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var newUser *dto2.NewUser
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
	var user *dto2.User
	var appErr *errs.AppError

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE.", err)
		return
	}

	user, err = dto2.Deserialize[dto2.User](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto2.Auth).AuthLvl == authLvl.ADMIN ||
		req.Context().Value("auth").(dto2.Auth).UserId == user.UUID {

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
	var user *dto2.User
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

	if req.Context().Value("auth").(dto2.Auth).UserId == id {

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

func (uh UserHandler) DeleteUser(resp http.ResponseWriter, req *http.Request) {
	var err *errs.AppError
	if req.Context().Value("auth").(dto2.Auth).AuthLvl == authLvl.ADMIN {
		params := mux.Vars(req)
		err = uh.UserService.DeleteUser(params["slug"])
	} else {
		err = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}
	if err != nil {
		WriteResponse(resp, err.Code, &dto2.User{}, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, &dto2.User{}, "")
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
