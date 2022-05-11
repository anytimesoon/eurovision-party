package main

import (
	initializer "eurovision/pkg/init"
	"fmt"
	"log"
)

func init() {
	db, err := initializer.Connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	err = initializer.CreateCountriesTable(db)
	if err != nil {
		log.Printf("Create country table failed with error %s", err)
		return
	}

	err = initializer.AddCountries(db)
	if err != nil {
		log.Printf("Adding countries failed with error %s", err)
		return
	}
}

func main() {

	fmt.Println("Done")
}
