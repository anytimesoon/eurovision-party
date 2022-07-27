package handler

import (
	"encoding/json"
	"eurovision/pkg/service"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	Service service.CommentService
}

func (ch CommentHandler) FindAllComments(resp http.ResponseWriter, req *http.Request) {
	comments, err := ch.Service.FindAllComments()
	if err != nil {
		log.Println("Error finding all comments", err)
	}

	json.NewEncoder(resp).Encode(comments)
}

func (ch CommentHandler) CreateComment(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COMMENT CREATE!", err)
		return
	}

	comment, err := ch.Service.CreateComment(body)
	if err != nil {
		log.Println("Failed to create comment", err)
		return
	}

	json.NewEncoder(resp).Encode(comment)
}

func (ch CommentHandler) RemoveComment(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := ch.Service.DeleteComment(params["uuid"])
	if err != nil {
		log.Println("FAILED to delete comment", err)
		return
	}

	json.NewEncoder(resp)
}
