package users

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/domain"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func All(writer http.ResponseWriter, req *http.Request) {
	users, err := dao.Users()
	if err != nil {
		log.Println("FAILED to find all users!")
		return
	}

	json.NewEncoder(writer).Encode(users)
}

func FindOne(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userSlug := params["slug"]

	partialUser := domain.User{Slug: userSlug}
	user, err := dao.User(partialUser)
	if err != nil {
		log.Printf("FAILED to find %s", userSlug)
	}

	json.NewEncoder(writer).Encode(user)
}

func Create(writer http.ResponseWriter, req *http.Request) {
	var receivedUser domain.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER CREATE!", err)
		return
	}

	err = json.Unmarshal(body, &receivedUser)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return
	}

	newUser, err := dao.CreateUser(receivedUser)
	if err != nil {
		log.Println("FAILED to create new user", err)
		return
	}

	json.NewEncoder(writer).Encode(newUser)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var receivedUser domain.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!", err)
		return
	}

	err = json.Unmarshal(body, &receivedUser)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return
	}

	user, err := dao.User(receivedUser)
	if err != nil {
		log.Printf("FAILED to find %s", receivedUser.Name)
	}

	updatedUser, err := dao.UsersUpdate(user, receivedUser)
	if err != nil {
		log.Printf("FAILED to update %s", receivedUser.Name)
	}

	json.NewEncoder(writer).Encode(updatedUser)
}

func RemoveUser(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userSlug := params["slug"]

	partialUser := domain.User{Slug: userSlug}
	user, err := dao.UserDelete(partialUser)
	if err != nil {
		log.Printf("FAILED to find %s", userSlug)
	}

	json.NewEncoder(writer).Encode(user)
}
