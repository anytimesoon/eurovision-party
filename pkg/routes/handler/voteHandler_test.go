package handler

import (
	"bytes"
	"encoding/json"
	"eurovision/mocks/service"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var voteRouter *mux.Router
var vh VoteHandler
var mockVoteService *service.MockVoteService

// var mockVotes []dto.Vote
var mockVote dto.Vote
var voteJSON []byte
var voteBody *bytes.Buffer

func setupVoteTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVoteService = service.NewMockVoteService(ctrl)
	vh = VoteHandler{mockVoteService}

	mockVote = dto.Vote{UUID: uuid.New(), UserId: uuid.New(), CountryId: uuid.New(), Costume: 2, Song: 4, Performance: 6, Props: 3}
	voteJSON, _ = json.Marshal(mockVote)
	voteBody = bytes.NewBuffer(voteJSON)

	voteRouter = mux.NewRouter()
	voteRouter.HandleFunc("/vote", vh.CreateVote).Methods(http.MethodPost)
	voteRouter.HandleFunc("/vote", vh.UpdateVote).Methods(http.MethodPut)
}

func Test_new_vote_route_returns_500_code(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().CreateVote(voteJSON).Return(nil, errs.NewUnexpectedError("Couldn't create new vote"))

	req, _ := http.NewRequest(http.MethodPost, "/vote", voteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_new_vote_route_returns_200_code(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().CreateVote(voteJSON).Return(&mockVote, nil)

	req, _ := http.NewRequest(http.MethodPost, "/vote", voteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}
