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
	countryRouter.HandleFunc("/", countryHandler.UpdateCountry).Methods(http.MethodPut)
	countryRouter.HandleFunc("/{slug}", countryHandler.FindOneCountry).Methods(http.MethodGet)

	// // User
	userRepositoryDb := domain.NewUserRepositoryDb(db)
	userHandler := handler.UserHandler{Service: service.NewUserService(userRepositoryDb)}
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/", userHandler.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/new", userHandler.CreateUser).Methods((http.MethodPost))
	userRouter.HandleFunc("/{slug}", userHandler.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}/rem", userHandler.RemoveUser).Methods(http.MethodDelete)

	// // Comment
	commentRepositoryDb := domain.NewCommentRepositoryDb(db)
	commentHandler := handler.CommentHandler{Service: service.NewCommentService(commentRepositoryDb)}
	commentRouter := router.PathPrefix("/comment").Subrouter()
	commentRouter.HandleFunc("/", commentHandler.FindAllComments).Methods(http.MethodGet)
	commentRouter.HandleFunc("/new", commentHandler.CreateComment).Methods((http.MethodPost))
	commentRouter.HandleFunc("/{uuid}/rem", commentHandler.RemoveComment).Methods(http.MethodDelete)

	// // Vote
	voteRepositoryDb := domain.NewVoteRepositoryDb(db)
	voteHandler := handler.VoteHandler{Service: service.NewVoteService(voteRepositoryDb)}
	voteRouter := router.PathPrefix("/vote").Subrouter()
	voteRouter.HandleFunc("/new", voteHandler.CreateVote).Methods(http.MethodPost)
	voteRouter.HandleFunc("/", voteHandler.UpdateVote).Methods(http.MethodPut)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))) //keeps the server alive on port 8080
}
