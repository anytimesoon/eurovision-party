package countries

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/domain"
	"io/ioutil"
	"log"
	"net/http"
)

func All(writer http.ResponseWriter, req *http.Request) {
	countries, err := dao.Countries()
	if err != nil {
		log.Println("home FAILED!")
		return
	}

	json.NewEncoder(writer).Encode(countries)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	var country domain.Country

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("FAILED to read body of COUNTRY UPDATE!")
		return
	}

	err = json.Unmarshal(body, &country)
	if err != nil {
		log.Println("country FAILED to UPDATE!")
		return
	}

	log.Printf("%+v", country)
}
