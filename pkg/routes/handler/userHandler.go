package handler

import (
	"encoding/json"
	"eurovision/pkg/service"
	"log"
	"net/http"
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

// func FindOneUser(writer http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	userSlug := params["slug"]

// 	userDAO, err := dao.SingleUser(userSlug)
// 	if err != nil {
// 		log.Printf("FAILED to find %s", userSlug)
// 	}

// 	userDTO := dto.User{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
// 	}

// 	json.NewEncoder(writer).Encode(userDTO)
// }

// func CreateUser(writer http.ResponseWriter, req *http.Request) {
// 	var userDTO dto.User

// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Println("FAILED to read body of USER CREATE!", err)
// 		return
// 	}

// 	err = json.Unmarshal(body, &userDTO)
// 	if err != nil {
// 		log.Println("FAILED to unmarshal json!", err)
// 		return
// 	}

// 	userDAO, err := dao.CreateUser(userDTO)
// 	if err != nil {
// 		log.Println("FAILED to create new user", err)
// 		return
// 	}

// 	userDTO = dto.User{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
// 	}

// 	json.NewEncoder(writer).Encode(userDTO)
// }

// func UpdateUser(writer http.ResponseWriter, req *http.Request) {
// 	var userDTO dto.User

// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Println("FAILED to read body of USER UPDATE!", err)
// 		return
// 	}

// 	err = json.Unmarshal(body, &userDTO)
// 	if err != nil {
// 		log.Println("FAILED to unmarshal json!", err)
// 		return
// 	}

// 	userDAO, err := dao.SingleUser(userDTO.Data.Slug)
// 	if err != nil {
// 		log.Printf("FAILED to find %s", userDTO.Data.Name)
// 	}

// 	userDAO, err = dao.UpdateUser(userDAO, userDTO)
// 	if err != nil {
// 		log.Printf("FAILED to update %s", userDTO.Data.Name)
// 	}

// 	userDTO = dto.User{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.UserData{UUID: userDAO.UUID, Name: userDAO.Name, Slug: userDAO.Slug, Icon: userDAO.Icon},
// 	}

// 	json.NewEncoder(writer).Encode(userDTO)
// }

// func RemoveUser(writer http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	userSlug := params["slug"]

// 	userDAO, err := dao.SingleUser(userSlug)
// 	if err != nil {
// 		log.Printf("FAILED to find %s", userSlug)
// 	}

// 	userName, err := dao.DeleteUser(userDAO)
// 	if err != nil {
// 		log.Printf("FAILED to delete %s with ERROR %s", userDAO.Name, err)
// 	}

// 	userDTO := dto.User{
// 		Success: true,
// 		Message: userName + " was deleted successfully",
// 	}

// 	json.NewEncoder(writer).Encode(userDTO)
// }
