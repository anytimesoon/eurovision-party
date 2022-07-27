package handler

import (
	"encoding/json"
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

	vote, err := vh.Service.CreateVote(body)
	if err != nil {
		log.Println("Failed to create vote", err)
		return
	}

	json.NewEncoder(resp).Encode(vote)
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	vote, err := vh.Service.UpdateVote(body)
	if err != nil {
		log.Println("Failed to update vote", err)
		return
	}

	json.NewEncoder(resp).Encode(vote)
}
