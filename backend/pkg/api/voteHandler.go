package api

import (
	"io"
	"log"
	"net/http"

	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/gorilla/mux"
)

type VoteHandler struct {
	Service service.VoteService
}

func (vh VoteHandler) UpdateVote(resp http.ResponseWriter, req *http.Request) {
	var appErr *errs.AppError
	var voteSingle *dto.VoteSingle
	vote := &dto.Vote{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of VOTE UPDATE!", err)
		return
	}

	voteSingle, err = dto.Deserialize[dto.VoteSingle](body)
	if err != nil {
		return
	}

	if req.Context().Value("auth").(dto.Auth).UserId == voteSingle.UserId {
		vote, appErr = vh.Service.UpdateVote(*voteSingle)
	} else {
		appErr = errs.NewUnauthorizedError(errs.Common.Unauthorized)
	}

	if appErr != nil {
		WriteResponse(resp, appErr.Code, *vote, appErr.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *vote, "")
	}
}

func (vh VoteHandler) GetVoteByUserAndCountry(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := req.Context().Value("auth").(dto.Auth).UserId

	vote, err := vh.Service.GetVoteByUserAndCountry(userId, params["slug"])

	if err != nil {
		WriteResponse(resp, err.Code, *vote, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *vote, "")
	}
}

func (vh VoteHandler) GetResults(resp http.ResponseWriter, req *http.Request) {
	results, err := vh.Service.GetResults()

	if err != nil {
		WriteResponse(resp, err.Code, *results, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *results, "")
	}
}

func (vh VoteHandler) GetResultsByUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	results, err := vh.Service.GetResultsByUser(params["userId"])

	if err != nil {
		WriteResponse(resp, err.Code, *results, err.Message)
	} else {
		WriteResponse(resp, http.StatusOK, *results, "")
	}
}
