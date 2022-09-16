package handler

import (
	"eurovision/pkg/service"
	"io/ioutil"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah AuthHandler) Register(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
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

}

func (ah AuthHandler) Authenticate(resp http.ResponseWriter, req *http.Request) {

}
