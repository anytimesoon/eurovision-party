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

func (db VoteRepositoryDb) CreateVote(voteDTO dto.Vote) (*Vote, *errs.AppError) {
	var vote Vote

	query := fmt.Sprintf(`INSERT INTO vote(uuid, userId, countrySlug, costume, song, performance, props) VALUES ('%s', '%s', '%s', %d, %d, %d, %d)`, voteDTO.UUID.String(), voteDTO.UserId, voteDTO.CountrySlug, voteDTO.Costume, voteDTO.Song, voteDTO.Performance, voteDTO.Props)

	_, err := db.client.NamedExec(query, vote)
	if err != nil {
		log.Println("Error when creating new vote:", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	query = fmt.Sprintf(`SELECT * FROM vote WHERE uuid = '%s'`, voteDTO.UUID.String())

	err = db.client.Get(&vote, query)
	if err != nil {
		log.Println("Error when fetching vote after create:", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your vote")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) UpdateVote(voteDTO dto.VoteSingle) (*Vote, *errs.AppError) {
	var vote Vote

	query := fmt.Sprintf(`UPDATE vote SET %s = %d WHERE userId = '%s'`, voteDTO.Cat, voteDTO.Score, voteDTO.UserId.String())

	_, err := db.client.NamedExec(query, vote)
	if err != nil {
		log.Println("Error while updating vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	query = fmt.Sprintf(`SELECT * FROM vote WHERE userId = '%s'`, voteDTO.UserId.String())
	err = db.client.Get(&vote, query)
	if err != nil {
		log.Println("Error while fetching vote after update", err)
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your vote")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*Vote, *errs.AppError) {
	var vote Vote

	query := fmt.Sprintf(`SELECT * FROM vote WHERE userId = '%s' AND countrySlug = '%s'`, userId.String(), countrySlug)
	err := db.client.Get(&vote, query)
	if err != nil && err.Error() == "sql: no rows in result set" {
		log.Println("Failed to find vote from country and user. Creating a new vote")

		return db.CreateVote(dto.Vote{
			UUID:        uuid.New(),
			UserId:      userId,
			CountrySlug: countrySlug,
			Costume:     0,
			Song:        0,
			Performance: 0,
			Props:       0,
		})
	} else if err != nil {
		log.Println("Failed to find vote from country and user.", err)
	}

	return &vote, nil
}
