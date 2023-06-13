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

	createVoteQuery := "INSERT INTO vote(userId, countrySlug, costume, song, performance, props) VALUES (?, ?, ?, ?, ?, ?)"
	getVoteQuery := "SELECT * FROM vote WHERE userId = ? AND countrySlug = ?"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error when starting new vote transaction. %s", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	_, err = tx.Exec(createVoteQuery, voteDTO.UserId.String(), voteDTO.CountrySlug, voteDTO.Costume, voteDTO.Song, voteDTO.Performance, voteDTO.Props)
	if err != nil {
		log.Printf("Error when creating new vote for user %s. %s", voteDTO.UserId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	err = tx.Get(&vote, getVoteQuery, voteDTO.UserId.String(), voteDTO.CountrySlug)
	if err != nil {
		log.Printf("Error when fetching vote for user %s after create. %s", voteDTO.UserId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error when committing new vote transaction. %s", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) UpdateVote(voteDTO dto.VoteSingle) (*Vote, *errs.AppError) {
	var vote Vote

	updateVoteQuery := fmt.Sprintf("UPDATE vote SET %s = ? WHERE userId = ? AND countrySlug = ?", voteDTO.Cat)
	getVoteQuery := "SELECT * FROM vote WHERE userId = ? AND countrySlug = ?"

	tx, err := db.client.Beginx()
	if err != nil {
		log.Printf("Error while starting transaction to update vote for user %s. %s", voteDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}
	_, err = tx.Exec(updateVoteQuery, voteDTO.Score, voteDTO.UserId.String(), voteDTO.CountrySlug)
	if err != nil {
		log.Printf("Error while updating vote for user %s. %s", voteDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	err = tx.Get(&vote, getVoteQuery, voteDTO.UserId.String(), voteDTO.CountrySlug)
	if err != nil {
		log.Printf("Error while fetching vote for user %s after update. %s", voteDTO.UserId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while committing transaction to update vote for user %s. %s", voteDTO.UserId, err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*Vote, *errs.AppError) {
	var vote Vote

	query := "SELECT * FROM vote WHERE userId = ? AND countrySlug = ?"
	err := db.client.Get(&vote, query, userId.String(), countrySlug)
	if err != nil && err.Error() == "sql: no rows in result set" {
		log.Println("Found 0 votes from country and user. Creating a new vote")

		return db.CreateVote(dto.Vote{
			UserId:      userId,
			CountrySlug: countrySlug,
			Costume:     0,
			Song:        0,
			Performance: 0,
			Props:       0,
		})
	} else if err != nil {
		log.Printf("Failed to find vote for country %s and user %s. %s", countrySlug, userId.String(), err)
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetAllVotes() (*[]Vote, *errs.AppError) {
	votes := make([]Vote, 0)

	query := "SELECT * FROM vote"
	err := db.client.Select(&votes, query)
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &votes, nil
}

func (db VoteRepositoryDb) GetResults() (*[]Result, *errs.AppError) {
	results := make([]Result, 0)

	query := `select countrySlug,
					   costume_total,
					   song_total,
					   performance_total,
					   props_total,
					   costume_total + song_total + performance_total + props_total as total
				from (select countrySlug, sum(costume) as costume_total,
							sum(song) as song_total,
							sum(performance) as performance_total,
							sum(props) as props_total
						from vote
						group by countrySlug) v
				order by total desc;`

	err := db.client.Select(&results, query)
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return &results, nil
}
