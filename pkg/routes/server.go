package routes

import (
	"eurovision/assets"
	"eurovision/conf"
	"eurovision/pkg/domain"
	"eurovision/pkg/routes/handler"
	"eurovision/pkg/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func StartServer(db *sqlx.DB, appConf conf.App) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logging)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(jsHeaders)

	// Country
	countryRepositoryDb := domain.NewCountryRepositoryDb(db)
	countryHandler := handler.CountryHandler{Service: service.NewCountryService(countryRepositoryDb)}
	countryRouter := apiRouter.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countryHandler.FindAllCountries).Methods(http.MethodGet)
	countryRouter.HandleFunc("/", countryHandler.UpdateCountry).Methods(http.MethodPut)
	countryRouter.HandleFunc("/{slug}", countryHandler.FindOneCountry).Methods(http.MethodGet)

	// User
	userRepositoryDb := domain.NewUserRepositoryDb(db)
	userHandler := handler.UserHandler{Service: service.NewUserService(userRepositoryDb)}
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/", userHandler.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/", userHandler.CreateUser).Methods((http.MethodPost))
	userRouter.HandleFunc("/{slug}", userHandler.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.RemoveUser).Methods(http.MethodDelete)

	// Vote
	voteRepositoryDb := domain.NewVoteRepositoryDb(db)
	voteHandler := handler.VoteHandler{Service: service.NewVoteService(voteRepositoryDb)}
	voteRouter := apiRouter.PathPrefix("/vote").Subrouter()
	voteRouter.HandleFunc("/", voteHandler.CreateVote).Methods(http.MethodPost)
	voteRouter.HandleFunc("/", voteHandler.UpdateVote).Methods(http.MethodPut)
	// voteRouter.HandleFunc("/user/{userId}", voteHandler.VoteByUser).Methods(http.MethodGet)
	// voteRouter.HandleFunc("/country/{countryId}", voteHandler.VoteByUser).Methods(http.MethodGet)

	// Assets
	fs := assets.NewStaticImageFS()
	imageRouter := router.PathPrefix("/img").Subrouter()
	imageRouter.PathPrefix("/static/").Handler(http.StripPrefix("/img/static/", fs)).Methods(http.MethodGet)
	imageRouter.Use(imgHeaders)

	// Chatroom
	commentRepositoryDb := domain.NewCommentRepositoryDb(db)
	commentService := service.NewCommentService(commentRepositoryDb)
	chatRoomHandler := handler.ChatRoomHandler{
		RoomService:    service.NewRoom(commentService),
		CommentService: commentService,
	}
	go chatRoomHandler.RoomService.Run()
	chatRouter := router.PathPrefix("/chat").Subrouter()
	chatRouter.HandleFunc("/connect", chatRoomHandler.Connect)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("Server listening on port %s", appConf.Server.Port)
	url := fmt.Sprintf("%s:%s", appConf.Server.Url, appConf.Server.Port)
	log.Fatal(http.ListenAndServe(url, handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
