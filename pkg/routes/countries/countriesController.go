package countries

import (
	"encoding/json"
	"eurovision/db"
	"eurovision/pkg/dao"
	dto "eurovision/pkg/dto"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func All(writer http.ResponseWriter, req *http.Request) {
	countriesDAO, err := dao.Countries(db.Conn)
	if err != nil {
		log.Println("home FAILED!")
		return
	}

	countriesDTO := dto.Countries{
		Success: true,
		Message: "",
	}

	for _, country := range countriesDAO {
		countriesDTO.Data = append(countriesDTO.Data, dto.CountryData{ID: country.UUID, Flag: country.Flag, Name: country.Name, Slug: country.Slug, BandName: country.BandName, SongName: country.SongName, Participating: country.Participating})
	}

	json.NewEncoder(writer).Encode(countriesDTO)
}

func FindOne(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	countrySlug := params["slug"]

	var countryDTO dto.Country
	countryDTO.Data.Slug = countrySlug

	country, err := dao.SingleCountry(countryDTO, db.Conn)
	if err != nil {
		log.Printf("FAILED to find %s", countrySlug)
	}

	countryDTO = dto.Country{
		Success: true,
		Message: "",
		Data:    dto.CountryData{ID: country.UUID, Flag: country.Flag, Name: country.Name, Slug: country.Slug, BandName: country.BandName, SongName: country.SongName, Participating: country.Participating},
	}

	json.NewEncoder(writer).Encode(countryDTO)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var countryDTO dto.Country

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COUNTRY UPDATE!")
		return
	}

	err = json.Unmarshal(body, &countryDTO.Data)
	if err != nil {
		log.Println("FAILED to unmarshal json!")
		return
	}

	countryDAO, err := dao.SingleCountry(countryDTO, db.Conn)
	if err != nil {
		log.Printf("FAILED to find %s", countryDTO.Data.Name)
	}

	countryDAO, err = dao.CountriesUpdate(countryDAO, countryDTO, db.Conn)
	if err != nil {
		log.Printf("FAILED to update %s", countryDTO.Data.Name)
	}

	countryDTO = dto.Country{
		Success: true,
		Message: "",
		Data:    dto.CountryData{ID: countryDAO.UUID, Flag: countryDAO.Flag, Name: countryDAO.Name, Slug: countryDAO.Slug, BandName: countryDAO.BandName, SongName: countryDAO.SongName, Participating: countryDAO.Participating},
	}

	json.NewEncoder(writer).Encode(countryDTO)
}
