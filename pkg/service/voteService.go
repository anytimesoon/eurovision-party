package service

import (
	dto2 "github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/google/uuid"
)

type VoteService interface {
	UpdateVote(dto2.VoteSingle) (*dto2.Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dto2.Vote, *errs.AppError)
	GetResults() (*[]dto2.Result, *errs.AppError)
	GetResultsByUser(userId string) (*[]dto2.Result, *errs.AppError)
}

type DefaultVoteService struct {
	repo data.VoteRepository
}

func NewVoteService(repo data.VoteRepository) DefaultVoteService {
	return DefaultVoteService{repo}
}

func (service DefaultVoteService) UpdateVote(voteSingleDTO dto2.VoteSingle) (*dto2.Vote, *errs.AppError) {
	appErr := voteSingleDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, appErr := service.repo.UpdateVote(voteSingleDTO)
	if appErr != nil {
		return nil, appErr
	}

	result := vote.ToDto()
	return &result, nil
}

func (service DefaultVoteService) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dto2.Vote, *errs.AppError) {
	vote, err := service.repo.GetVoteByUserAndCountry(userId, countrySlug)
	if err != nil {
		return nil, err
	}

	result := vote.ToDto()
	return &result, nil
}

func (service DefaultVoteService) GetResults() (*[]dto2.Result, *errs.AppError) {
	results, err := service.repo.GetResults()
	if err != nil {
		return nil, err
	}

	resultsDTO := make([]dto2.Result, 0)
	for _, result := range *results {
		resultDTO := result.ToDto()
		resultsDTO = append(resultsDTO, resultDTO)
	}

	return &resultsDTO, nil
}

func (service DefaultVoteService) GetResultsByUser(userId string) (*[]dto2.Result, *errs.AppError) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	results, appErr := service.repo.GetResultsByUser(id)
	if appErr != nil {
		return nil, appErr
	}

	resultsDTO := make([]dto2.Result, 0)
	for _, result := range *results {
		resultDTO := result.ToDto()
		resultsDTO = append(resultsDTO, resultDTO)
	}

	return &resultsDTO, nil
}
