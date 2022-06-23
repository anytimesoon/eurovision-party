package main

import (
	"database/sql"
	"encoding/json"
	dao "eurovision/pkg/dao"
	"eurovision/pkg/domain"
	initializer "eurovision/pkg/init"
	"io/ioutil"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/country", updateCountry).Methods(http.MethodPost)
	router.Use(addHeaders, logging)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))) //keeps the server alive on port 8080

	db.Close()
}

func addHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s was requested by %q", r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func home(writer http.ResponseWriter, req *http.Request) {
	countries, err := dao.Countries(db)
	if err != nil {
		log.Println("home FAILED!")
		return
	}

	json.NewEncoder(writer).Encode(countries)
}

func updateCountry(writer http.ResponseWriter, req *http.Request) {
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
