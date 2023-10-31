package router

import (
	"bytes"
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/mocks/service"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
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

var mockVote dto.Vote
var voteJSON []byte
var voteBody *bytes.Buffer
var invalidVote dto.Vote
var invalidVoteJSON []byte
var invalidVoteBody *bytes.Buffer

func setupVoteTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVoteService = service.NewMockVoteService(ctrl)
	vh = VoteHandler{mockVoteService}

	mockVote = dto.Vote{UUID: uuid.New(), UserId: uuid.New(), CountryId: uuid.New(), Costume: 2, Song: 4, Performance: 6, Props: 3}
	voteJSON, _ = json.Marshal(mockVote)
	voteBody = bytes.NewBuffer(voteJSON)

	invalidVote = dto.Vote{UUID: uuid.New(), UserId: uuid.New(), CountryId: uuid.New(), Costume: 100, Song: 4, Performance: 6, Props: 3}
	invalidVoteJSON, _ = json.Marshal(mockVote)
	invalidVoteBody = bytes.NewBuffer(voteJSON)

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

func Test_new_vote_route_returns_400_error(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().CreateVote(invalidVoteJSON).Return(nil, errs.NewInvalidError("Vote name must not be blank"))

	req, _ := http.NewRequest(http.MethodPost, "/vote", invalidVoteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", recorder.Code)
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

func Test_vote_update_route_returns_500_code(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().UpdateVote(voteJSON).Return(nil, errs.NewUnexpectedError("Couldn't update vote"))

	req, _ := http.NewRequest(http.MethodPut, "/vote", voteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Expected status code 500, but got", recorder.Code)
	}
}

func Test_vote_update_route_returns_400_error(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().UpdateVote(invalidVoteJSON).Return(nil, errs.NewInvalidError("Vote must be less than 5"))

	req, _ := http.NewRequest(http.MethodPut, "/vote", invalidVoteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", recorder.Code)
	}
}

func Test_vote_update_route_returns_200_code(t *testing.T) {
	setupVoteTest(t)

	mockVoteService.EXPECT().UpdateVote(voteJSON).Return(&mockVote, nil)

	req, _ := http.NewRequest(http.MethodPut, "/vote", voteBody)

	recorder := httptest.NewRecorder()
	voteRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expected status code 200, but got", recorder.Code)
	}
}
