package service

import (
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type VoteService interface {
	UpdateVote(dto.VoteSingle) (*dto.Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dto.Vote, *errs.AppError)
	GetResults() (*[]dto.Result, *errs.AppError)
	GetResultsByUser(userId string) (*[]dto.Result, *errs.AppError)
}

type DefaultVoteService struct {
	repo data.VoteRepository
}

func NewVoteService(repo data.VoteRepository) DefaultVoteService {
	return DefaultVoteService{repo}
}

func (vs DefaultVoteService) UpdateVote(voteSingleDTO dto.VoteSingle) (*dto.Vote, *errs.AppError) {
	appErr := voteSingleDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, err := vs.repo.GetVoteByUserAndCountry(voteSingleDTO.UserId, voteSingleDTO.CountrySlug)
	if err != nil {
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your votes")
	}

	switch voteSingleDTO.Cat {
	case enum.Song:
		vote.Song = voteSingleDTO.Score
	case enum.Costume:
		vote.Costume = voteSingleDTO.Score
	case enum.Performance:
		vote.Performance = voteSingleDTO.Score
	case enum.Props:
		vote.Props = voteSingleDTO.Score
	}

	vote.Total = int(vote.Costume) + int(vote.Song) + int(vote.Performance) + int(vote.Props)

	vote, err = vs.repo.UpdateVote(*vote)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	result := vote.ToDto()
	return &result, nil
}

func (vs DefaultVoteService) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dto.Vote, *errs.AppError) {
	vote, err := vs.repo.GetVoteByUserAndCountry(userId, countrySlug)
	if err != nil {
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your votes")
	}

	result := vote.ToDto()
	return &result, nil
}

func (vs DefaultVoteService) GetResults() (*[]dto.Result, *errs.AppError) {
	results, err := vs.repo.GetResults()
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return sortResults(*results), nil
}

func (vs DefaultVoteService) GetResultsByUser(userId string) (*[]dto.Result, *errs.AppError) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	results, err := vs.repo.GetResultsByUser(id)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return sortResults(*results), nil
}

func sortResults(results []dao.Result) *[]dto.Result {
	sortedResultsDTO := make([]dto.Result, 0)
	for _, result := range results {
		resultDTO := result.ToDto()

		inserted := false
		for i := range sortedResultsDTO {
			if resultDTO.Total > sortedResultsDTO[i].Total {
				sortedResultsDTO = append(sortedResultsDTO[:i+1], sortedResultsDTO[i:]...)
				sortedResultsDTO[i] = resultDTO
				inserted = true
				break
			}
		}

		if !inserted {
			sortedResultsDTO = append(sortedResultsDTO, resultDTO)
		}
	}

	return &sortedResultsDTO
}
