package domain

import (
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VoteRepositoryDb struct {
	client *sqlx.DB
}

func NewVoteRepositoryDb(db *sqlx.DB) VoteRepositoryDb {
	return VoteRepositoryDb{db}
}

func (db VoteRepositoryDb) CreateVote(voteDTO *dto.Vote) (*Vote, *errs.AppError) {
	var vote Vote

	uuid := uuid.New().String()

	query := fmt.Sprintf(`INSERT INTO vote(uuid, userId, countryId, costume, song, performance, props) VALUES ('%s', '%s', '%s', %d, %d, %d, %d)`, uuid, voteDTO.UserId, voteDTO.CountryId, voteDTO.Costume, voteDTO.Song, voteDTO.Performance, voteDTO.Props)

	_, err := db.client.NamedExec(query, vote)
	if err != nil {
		log.Println("Error when creating new vote:", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	query = fmt.Sprintf(`SELECT * FROM vote WHERE uuid = '%s'`, uuid)

	err = db.client.Get(&vote, query)
	if err != nil {
		log.Println("Error when fetching vote after create:", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your vote")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) UpdateVote(voteDTO *dto.Vote) (*Vote, *errs.AppError) {
	var vote Vote

	query := fmt.Sprintf(`UPDATE vote SET costume = %d, song = %d, performance = %d, props = %d WHERE uuid = '%s'`, voteDTO.Costume, voteDTO.Song, voteDTO.Performance, voteDTO.Props, voteDTO.UUID.String())

	_, err := db.client.NamedExec(query, vote)
	if err != nil {
		log.Println("Error while updating vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	query = fmt.Sprintf(`SELECT * FROM vote WHERE uuid = '%s'`, voteDTO.UUID.String())
	err = db.client.Get(&vote, query)
	if err != nil {
		log.Println("Error while fetching vote after update", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your vote")
	}

	return &vote, nil
}
