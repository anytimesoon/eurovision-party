package main

import (
	"database/sql"
	"encoding/json"
	dao "eurovision/pkg/dao"
	initializer "eurovision/pkg/init"
	"fmt"
	"log"
	"mime"
	"net/http"
)

var db *sql.DB

func init() {
	sqlDb, err := initializer.Connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}

	log.Printf("Successfully connected to database")

	err = initializer.CreateCountriesTable(sqlDb)
	if err != nil {
		log.Printf("Create country table failed with error %s", err)
		return
	}

	err = initializer.AddCountries(sqlDb)
	if err != nil {
		log.Printf("Adding countries failed with error %s", err)
		return
	}

	db = sqlDb
}

func main() {
	mime.AddExtensionType(".js", "application/javascript")

	http.Handle("/", http.HandlerFunc(home))
	log.Fatal(http.ListenAndServe(":8080", nil))

	db.Close()
}

func home(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)

	countries, err := dao.Countries(db)
	if err != nil {
		fmt.Println("home FAILED!")
		return
	}

	json.NewEncoder(writer).Encode(countries)
}
