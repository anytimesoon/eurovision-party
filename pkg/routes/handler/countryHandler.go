package handler

import (
	"encoding/json"
	"eurovision/pkg/service"
	"net/http"
)

type CountryHandler struct {
	Service service.CountryService
}

func (ch *CountryHandler) FindAllCountries(wr http.ResponseWriter, req *http.Request) {
	countries, err := ch.Service.GetAllCountries()
	if err != nil {
		wr.WriteHeader(http.StatusNotFound)
		json.NewEncoder(wr).Encode(countries)
	} else {
		json.NewEncoder(wr).Encode(countries)
	}

}

// func FindOneCountry(writer http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	countrySlug := params["slug"]

// 	var countryDTO dto.Country
// 	countryDTO.Data.Slug = countrySlug

// 	country, err := dao.SingleCountry(countryDTO, db.Conn)
// 	if err != nil {
// 		log.Printf("FAILED to find %s", countrySlug)
// 	}

// 	countryDTO = dto.Country{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.CountryData{ID: country.UUID, Flag: country.Flag, Name: country.Name, Slug: country.Slug, BandName: country.BandName, SongName: country.SongName, Participating: country.Participating},
// 	}

// 	json.NewEncoder(writer).Encode(countryDTO)
// }

// func UpdateCountry(writer http.ResponseWriter, req *http.Request) {
// 	var countryDTO dto.Country

// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Println("FAILED to read body of COUNTRY UPDATE!")
// 		return
// 	}

// 	err = json.Unmarshal(body, &countryDTO.Data)
// 	if err != nil {
// 		log.Println("FAILED to unmarshal json!")
// 		return
// 	}

// 	countryDAO, err := dao.SingleCountry(countryDTO, db.Conn)
// 	if err != nil {
// 		log.Printf("FAILED to find %s", countryDTO.Data.Name)
// 	}

// 	countryDAO, err = dao.CountriesUpdate(countryDAO, countryDTO, db.Conn)
// 	if err != nil {
// 		log.Printf("FAILED to update %s", countryDTO.Data.Name)
// 	}

// 	countryDTO = dto.Country{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.CountryData{ID: countryDAO.UUID, Flag: countryDAO.Flag, Name: countryDAO.Name, Slug: countryDAO.Slug, BandName: countryDAO.BandName, SongName: countryDAO.SongName, Participating: countryDAO.Participating},
// 	}

// 	json.NewEncoder(writer).Encode(countryDTO)
// }
