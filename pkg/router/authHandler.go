package router

import (
	"eurovision/pkg/service"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Register(resp http.ResponseWriter, req *http.Request) {
	ok, appErr := currentSessions.authorize(req)
	if appErr != nil || !ok {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	auth, appErr := ah.Service.Register(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, auth)
	}
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	auth, err := ah.Service.Login(params["token"], params["userId"])
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    auth.Token,
			Expires:  auth.Expiration,
			Path:     "/",
			SameSite: 3,
		}
		http.SetCookie(resp, cookie)
		currentSessions.sessions[auth.Token] = session{
			userId:  auth.UserId,
			exp:     auth.Expiration,
			authLvl: auth.AuthLvl,
			slug:    auth.Slug,
		}
		writeResponse(resp, http.StatusOK, auth.ToDTO())
	}
}
