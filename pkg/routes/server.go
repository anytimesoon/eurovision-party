package routes

import (
	"eurovision/pkg/routes/countries"
	"eurovision/pkg/routes/users"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	mime.AddExtensionType(".js", "application/javascript")

	router := mux.NewRouter()

	// Country
	router.HandleFunc("/", countries.All).Methods(http.MethodGet)
	router.HandleFunc("/country", countries.Update).Methods(http.MethodPut)
	router.HandleFunc("/country/{slug}", countries.FindOne).Methods(http.MethodGet)

	// User
	router.HandleFunc("/user", users.All).Methods(http.MethodGet)
	router.HandleFunc("/user", users.Update).Methods(http.MethodPut)
	router.HandleFunc("/user/{slug}", users.FindOne).Methods(http.MethodGet)

	router.Use(addHeaders, logging)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))) //keeps the server alive on port 8080
}
