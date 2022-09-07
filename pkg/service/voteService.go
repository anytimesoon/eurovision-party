package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

//go:generate mockgen -source=voteService.go -destination=../../mocks/service/mockVoteService.go -package=service eurovision/pkg/service
type VoteService interface {
	CreateVote([]byte) (*dto.Vote, *errs.AppError)
	UpdateVote([]byte) (*dto.Vote, *errs.AppError)
}

type DefaultVoteService struct {
	repo domain.VoteRepository
}

func NewVoteService(repo domain.VoteRepository) DefaultVoteService {
	return DefaultVoteService{repo}
}

func (service DefaultVoteService) CreateVote(body []byte) (*dto.Vote, *errs.AppError) {
	var voteDTO dto.Vote
	err := json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	appErr := voteDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, appErr := service.repo.CreateVote(voteDTO)
	if appErr != nil {
		return nil, appErr
	}

	result := vote.ToDto()
	return &result, nil
}

func (service DefaultVoteService) UpdateVote(body []byte) (*dto.Vote, *errs.AppError) {
	var voteDTO dto.Vote
	err := json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	appErr := voteDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, appErr := service.repo.UpdateVote(voteDTO)
	if appErr != nil {
		return nil, appErr
	}

	result := vote.ToDto()
	return &result, nil
}
