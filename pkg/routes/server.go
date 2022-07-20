package routes

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/routes/handler"
	"eurovision/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func StartServer(db *sqlx.DB) {
	router := mux.NewRouter()
	router.Use(addHeaders, logging)

	// // Country
	countryRepositoryDb := domain.NewCountryRepositoryDb(db)
	countryHandler := handler.CountryHandler{Service: service.NewCountryService(countryRepositoryDb)}

	countryRouter := router.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countryHandler.FindAllCountries).Methods(http.MethodGet)
	// countryRouter.HandleFunc("/", countries.Update).Methods(http.MethodPut)
	// countryRouter.HandleFunc("/{slug}", countries.FindOne).Methods(http.MethodGet)

	// // User
	// userRouter := router.PathPrefix("/user").Subrouter()
	// userRouter.HandleFunc("/", users.All).Methods(http.MethodGet)
	// userRouter.HandleFunc("/", users.Update).Methods(http.MethodPut)
	// userRouter.HandleFunc("/new", users.Create).Methods((http.MethodPost))
	// userRouter.HandleFunc("/{slug}", users.FindOne).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{slug}/rem", users.RemoveUser).Methods(http.MethodDelete)

	// // Comment
	// commentRouter := router.PathPrefix("/comment").Subrouter()
	// commentRouter.HandleFunc("/", comments.All).Methods(http.MethodGet)
	// commentRouter.HandleFunc("/new", comments.Create).Methods((http.MethodPost))
	// // commentRouter.HandleFunc("/{uuid}/rem", comments.RemoveComment).Methods(http.MethodDelete)

	// // Vote
	// voteRouter := router.PathPrefix("/vote").Subrouter()
	// voteRouter.HandleFunc("/new", votes.Create).Methods(http.MethodPost)
	// voteRouter.HandleFunc("/", votes.Update).Methods(http.MethodPut)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))) //keeps the server alive on port 8080
}
