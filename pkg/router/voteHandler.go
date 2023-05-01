package router

import (
	"eurovision/pkg/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type VoteHandler struct {
	Service service.VoteService
}

func (vh VoteHandler) CreateVote(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE CREATE!", err)
		return
	}

	vote, appErr := vh.Service.CreateVote(body)
	if appErr != nil {
		writeResponse(resp, req, appErr.Code, vote, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, vote, "")
	}
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	vote, appErr := vh.Service.UpdateVote(body)
	if appErr != nil {
		writeResponse(resp, req, appErr.Code, vote, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, vote, "")
	}
}

func (vh VoteHandler) GetVoteByUserAndCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

}
