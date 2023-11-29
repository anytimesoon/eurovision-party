package service

import (
	"encoding/json"
	mockDomain "github.com/anytimesoon/eurovision-party/mocks/domain"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var voteService VoteService
var mockVoteRepository *mockDomain.MockVoteRepository
var mockVote domain.Vote
var mockVoteDTO dto.Vote
var voteJSON []byte
var invalidVoteDTO dto.Vote
var invalidVoteJSON []byte

func setupVoteTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVoteRepository = mockDomain.NewMockVoteRepository(ctrl)
	voteService = DefaultVoteService{mockVoteRepository}
	mockVote = domain.Vote{UUID: uuid.New(), UserId: uuid.New(), CountryId: uuid.New(), Costume: 2, Song: 4, Performance: 6, Props: 3}
	mockVoteDTO = mockVote.ToDto()
	voteJSON, _ = json.Marshal(mockVoteDTO)

	invalidVoteDTO = dto.Vote{UUID: uuid.New(), UserId: uuid.New(), CountryId: uuid.New(), Costume: 100, Song: 4, Performance: 6, Props: 3}
	invalidVoteJSON, _ = json.Marshal(invalidVoteDTO)
}

func Test_update_vote_service_returns_updated_vote(t *testing.T) {
	setupVoteTest(t)

	mockVoteRepository.EXPECT().UpdateVote(mockVoteDTO).Return(&mockVote, nil)

	result, _ := voteService.UpdateVote(voteJSON)

	if result.UUID != mockVoteDTO.UUID {
		t.Error("Returned votes do not match expected")
	}
}

func Test_update_vote_service_returns_500_error(t *testing.T) {
	setupVoteTest(t)

	mockVoteRepository.EXPECT().UpdateVote(mockVoteDTO).Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := voteService.UpdateVote(voteJSON)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_update_vote_service_returns_400_error(t *testing.T) {
	setupVoteTest(t)
	_, err := voteService.UpdateVote(invalidVoteJSON)

	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", err.Code)
	}
}

func Test_create_vote_service_returns_new_vote(t *testing.T) {
	setupVoteTest(t)

	mockVoteRepository.EXPECT().CreateVote(mockVoteDTO).Return(&mockVote, nil)

	result, _ := voteService.CreateVote(voteJSON)

	if result.UUID != mockVoteDTO.UUID {
		t.Error("Returned votes do not match expected")
	}
}

func Test_create_vote_service_returns_500_error(t *testing.T) {
	setupVoteTest(t)

	mockVoteRepository.EXPECT().CreateVote(mockVoteDTO).Return(nil, errs.NewUnexpectedError("DB error occurred"))

	_, err := voteService.CreateVote(voteJSON)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 error, but got %d", err.Code)
	}
}

func Test_create_vote_service_returns_400_error(t *testing.T) {
	setupVoteTest(t)
	_, err := voteService.UpdateVote(invalidVoteJSON)

	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 error, but got %d", err.Code)
	}
}
