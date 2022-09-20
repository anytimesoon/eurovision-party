package router

import (
	"eurovision/pkg/domain"
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
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, users)
	}
}

func (uh UserHandler) UpdateUser(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!", err)
		return
	}

	user, appErr := uh.Service.UpdateUser(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, user)
	}
}

func (uh UserHandler) FindOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.Service.SingleUser(params["slug"])
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, user)
	}
}

func (uh UserHandler) RemoveUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := uh.Service.DeleteUser(params["slug"])
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, domain.User{})
	}
}
