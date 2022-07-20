package dto

import (
	"github.com/google/uuid"
)

type CountryData struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	BandName      string    `json:"bandName"`
	SongName      string    `json:"songName"`
	Flag          string    `json:"flag"`
	Participating bool      `json:"participating"`
}

type Countries struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []CountryData `json:"data"`
}

type Country struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    CountryData `json:"data"`
}

// func SingleCountry(countryDTO dto.Country, dbc *sql.DB) (Country, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelFunc()
// 	var countryDAO Country

// 	query := fmt.Sprintf(`SELECT * FROM country WHERE slug = '%s'`, countryDTO.Data.Slug)
// 	stmt, err := dbc.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return countryDAO, err
// 	}

// 	row := stmt.QueryRowContext(ctx)

// 	err = row.Scan(&countryDAO.UUID, &countryDAO.Name, &countryDAO.Slug, &countryDAO.BandName, &countryDAO.SongName, &countryDAO.Flag, &countryDAO.Participating)
// 	if err != nil {
// 		log.Printf("FAILED to scan because %s", err)
// 		return countryDAO, err
// 	}

// 	return countryDAO, nil
// }

// func CountriesUpdate(countryDAO Country, countryDTO dto.Country, dbc *sql.DB) (Country, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelFunc()

// 	query := fmt.Sprintf(`UPDATE country SET bandName = '%s', songName = '%s', participating = %t WHERE uuid = '%s'`, countryDTO.Data.BandName, countryDTO.Data.SongName, countryDTO.Data.Participating, countryDTO.Data.ID.String())
// 	stmt, err := dbc.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return countryDAO, err
// 	}

// 	res, err := stmt.ExecContext(ctx)
// 	if err != nil {
// 		log.Printf("sql execution FAILED! %s was not updated %s", countryDAO.Name, err)
// 		return countryDAO, err
// 	}

// 	rowsAffected, err := res.RowsAffected()
// 	if err != nil || rowsAffected == 0 {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return countryDAO, err
// 	}

// 	countryDAO, err = SingleCountry(countryDTO, dbc)
// 	if err != nil {
// 		log.Printf("FAILED to find updated country %s", err)
// 		return countryDAO, err
// 	}

// 	log.Println("Country rows affected:", rowsAffected)
// 	return countryDAO, nil
// }
