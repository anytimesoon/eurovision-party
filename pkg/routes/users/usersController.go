package users

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/dto"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func All(writer http.ResponseWriter, req *http.Request) {
	usersDAO, err := dao.Users()
	if err != nil {
		log.Println("FAILED to find all users!")
		return
	}

	usersDTO := dto.Users{
		Success: true,
		Message: "",
	}

	for _, user := range usersDAO {
		usersDTO.Data = append(usersDTO.Data, dto.UserData{UUID: user.UUID, Name: user.Name, Slug: user.Slug, Icon: user.Icon})
	}

	json.NewEncoder(writer).Encode(usersDTO)
}

func FindOne(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userSlug := params["slug"]

	userDAO, err := dao.SingleUser(userSlug)
	if err != nil {
		log.Printf("FAILED to find %s", userSlug)
	}

	userDTO := dto.User{
		Success: true,
		Message: "",
		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
	}

	json.NewEncoder(writer).Encode(userDTO)
}

func Create(writer http.ResponseWriter, req *http.Request) {
	var userDTO dto.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER CREATE!", err)
		return
	}

	err = json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return
	}

	userDAO, err := dao.CreateUser(userDTO)
	if err != nil {
		log.Println("FAILED to create new user", err)
		return
	}

	userDTO = dto.User{
		Success: true,
		Message: "",
		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
	}

	json.NewEncoder(writer).Encode(userDTO)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var userDTO dto.User

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of USER UPDATE!", err)
		return
	}

	err = json.Unmarshal(body, &userDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return
	}

	userDAO, err := dao.SingleUser(userDTO.Data.Slug)
	if err != nil {
		log.Printf("FAILED to find %s", userDTO.Data.Name)
	}

	userDAO, err = dao.UpdateUser(userDAO, userDTO)
	if err != nil {
		log.Printf("FAILED to update %s", userDTO.Data.Name)
	}

	userDTO = dto.User{
		Success: true,
		Message: "",
		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
	}

	json.NewEncoder(writer).Encode(userDTO)
}

func RemoveUser(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userSlug := params["slug"]

	userDAO, err := dao.SingleUser(userSlug)
	if err != nil {
		log.Printf("FAILED to find %s", userSlug)
	}

	userName, err := dao.DeleteUser(userDAO)
	if err != nil {
		log.Printf("FAILED to delete %s with ERROR %s", userDAO.Name, err)
	}

	userDTO := dto.User{
		Success: true,
		Message: userName + " was deleted successfully",
	}

	json.NewEncoder(writer).Encode(userDTO)
}
