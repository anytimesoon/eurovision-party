package routes

import (
	"eurovision/pkg/routes/countries"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	mime.AddExtensionType(".js", "application/javascript")

	router := mux.NewRouter()
	router.HandleFunc("/", countries.All).Methods(http.MethodGet)
	router.HandleFunc("/country", countries.Update).Methods(http.MethodPut)
	router.HandleFunc("/country/{name}", countries.FindOne).Methods(http.MethodGet)
	router.Use(addHeaders, logging)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))) //keeps the server alive on port 8080
}
