package router

import (
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"io"
	"log"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Register(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var auth *dto.NewUser
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.ADMIN {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		auth, appErr = ah.Service.Register(body)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, *auth, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *auth, "")
	}
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of login request!", err)
		return
	}

	authorizedUser, appErr := ah.Service.Login(body)

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, &dto.SessionAuth{}, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, authorizedUser, "")
	}
}
