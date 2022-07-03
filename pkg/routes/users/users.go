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
	userName := params["name"]

	partialUser := domain.User{Name: userName}
	user, err := dao.User(partialUser)
	if err != nil {
		log.Printf("FAILED to find %s", userName)
	}

	json.NewEncoder(writer).Encode(user)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var receivedUser domain.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!")
		return
	}

	err = json.Unmarshal(body, &receivedUser)
	if err != nil {
		log.Println("FAILED to unmarshal json!")
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
