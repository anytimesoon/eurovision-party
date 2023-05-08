package router

import (
	"context"
	"eurovision/pkg/dto"
	"eurovision/pkg/enum"
	"eurovision/pkg/errs"
	"eurovision/pkg/service"
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
	if req.Context().Value("auth").(dto.Auth).AuthLvl == enum.Admin {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		auth, appErr = ah.Service.Register(body)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, auth, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, auth, "")
	}
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of login request!", err)
		return
	}

	auth, user, appErr := ah.Service.Login(body)
	if appErr != nil {

		ctx := context.WithValue(req.Context(), "auth", dto.Auth{})
		writeResponse(resp, req.WithContext(ctx), appErr.Code, user, appErr.Message)
	} else {
		ctx := context.WithValue(req.Context(), "auth", *auth)
		writeResponse(resp, req.WithContext(ctx), http.StatusOK, user, "")
	}
}
