package handler

import (
	"eurovision/pkg/domain"
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
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, comments)
	}
}

func (ch CommentHandler) CreateComment(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COMMENT CREATE!", err)
		return
	}

	comment, appErr := ch.Service.CreateComment(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, comment)
	}
}

func (ch CommentHandler) RemoveComment(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := ch.Service.DeleteComment(params["uuid"])
	if err != nil {
		writeResponse(resp, err.Code, err.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, domain.Comment{})
	}
}
