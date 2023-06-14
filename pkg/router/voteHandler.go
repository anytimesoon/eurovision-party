package router

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"eurovision/pkg/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type VoteHandler struct {
	Service service.VoteService
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var voteSingle *dto.VoteSingle
	var vote *dto.Vote

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	voteSingle, err = dto.Decode[dto.VoteSingle](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).UserId == voteSingle.UserId {
		vote, appErr = vh.Service.UpdateVote(*voteSingle)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, *vote, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *vote, "")
	}
}

func (vh VoteHandler) GetVoteByUserAndCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := req.Context().Value("auth").(dto.Auth).UserId

	vote, err := vh.Service.GetVoteByUserAndCountry(userId, params["slug"])

	if err != nil {
		writeResponse(resp, req, err.Code, *vote, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *vote, "")
	}
}

func (vh VoteHandler) GetResults(resp http.ResponseWriter, req *http.Request) {
	results, err := vh.Service.GetResults()

	if err != nil {
		writeResponse(resp, req, err.Code, *results, err.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, *results, "")
	}
}
