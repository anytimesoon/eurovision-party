package handler

import (
	"eurovision/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
)

type VoteHandler struct {
	Service service.VoteService
}

func (vh VoteHandler) CreateVote(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE CREATE!", err)
		return
	}

	vote, appErr := vh.Service.CreateVote(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, vote)
	}
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	vote, appErr := vh.Service.UpdateVote(body)
	if appErr != nil {
		writeResponse(resp, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(resp, http.StatusOK, vote)
	}
}
