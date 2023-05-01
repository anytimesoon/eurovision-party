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

func (vh VoteHandler) CreateVote(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var vote *dto.Vote

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE CREATE!", err)
		return
	}

	vote, err = dto.Decode[dto.Vote](body)
	if err != nil {
		return
	}

	if req.Context().Value("authAndToken").(dto.AuthAndToken).UUID == vote.UserId {
		vote, appErr = vh.Service.CreateVote(*vote)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, vote, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, vote, "")
	}
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var vote *dto.Vote

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	vote, err = dto.Decode[dto.Vote](body)
	if err != nil {
		return
	}

	if req.Context().Value("authAndToken").(dto.AuthAndToken).UUID == vote.UserId {
		vote, appErr = vh.Service.UpdateVote(*vote)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		writeResponse(resp, req, appErr.Code, vote, appErr.Message)
	} else {
		writeResponse(resp, req, http.StatusOK, vote, "")
	}
}

func (vh VoteHandler) GetVoteByUserAndCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

}
