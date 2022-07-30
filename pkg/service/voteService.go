package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"log"
)

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
	voteDTO, err := unmarshalVote(body)
	if err != nil {
		return nil, err
	}

	appErr := voteDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, AppError := service.repo.CreateVote(voteDTO)
	if AppError != nil {
		return nil, AppError
	}

	result := vote.ToDto()
	return &result, nil
}

func (service DefaultVoteService) UpdateVote(body []byte) (*dto.Vote, *errs.AppError) {
	voteDTO, err := unmarshalVote(body)
	if err != nil {
		return nil, err
	}

	appErr := voteDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, err := service.repo.UpdateVote(voteDTO)
	if err != nil {
		return nil, err
	}

	result := vote.ToDto()
	return &result, nil
}

func unmarshalVote(body []byte) (*dto.Vote, *errs.AppError) {
	var voteDTO dto.Vote
	err := json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}
	return &voteDTO, nil
}
