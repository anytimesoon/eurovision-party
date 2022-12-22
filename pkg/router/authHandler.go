package router

import (
	"eurovision/pkg/service"
	"io"
	"log"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Register(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	auth, appErr := ah.Service.Register(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, auth, appErr.Message)
	} else {
		writeResponse(resp, http.StatusOK, auth, "")
	}
}

func (ah AuthHandler) Login(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of login request!", err)
		return
	}

	auth, appErr := ah.Service.Login(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, auth, appErr.Message)
	} else {
		resp.Header().Add("Authorization", auth.EToken)
		writeResponse(resp, http.StatusOK, auth, "")
	}
}
