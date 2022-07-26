package handler

import (
	"encoding/json"
	"eurovision/pkg/service"
	"io/ioutil"
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
		log.Println("Failed to find all users", err)
	}

	json.NewEncoder(resp).Encode(users)
}

func (uh UserHandler) UpdateUser(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!", err)
		return
	}

	user, err := uh.Service.UpdateUser(body)
	if err != nil {
		log.Println("Failed to update user", err)
		return
	}

	json.NewEncoder(resp).Encode(user)
}

func (uh UserHandler) CreateUser(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER CREATE!", err)
		return
	}

	user, err := uh.Service.CreateUser(body)
	if err != nil {
		log.Println("Failed to create user", err)
		return
	}

	json.NewEncoder(resp).Encode(user)
}

func (uh UserHandler) FindOneUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user, err := uh.Service.SingleUser(params["slug"])
	if err != nil {
		log.Println("FAILED to find user", err)
		return
	}

	json.NewEncoder(resp).Encode(user)
}

func (uh UserHandler) RemoveUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := uh.Service.DeleteUser(params["slug"])
	if err != nil {
		log.Println("FAILED to delete user", err)
		return
	}

	json.NewEncoder(resp)
}
