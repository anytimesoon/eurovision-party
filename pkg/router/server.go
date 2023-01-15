package router

import (
	"eurovision/assets"
	"eurovision/conf"
	"eurovision/pkg/domain"
	"eurovision/pkg/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var authService service.AuthService

func StartServer(db *sqlx.DB, appConf conf.App) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logging)

	// Authentication
	authRepositoryMem := domain.NewAuthRepositoryDB(db)
	authService = service.NewAuthService(authRepositoryMem)
	authHandler := AuthHandler{Service: authService}
	router.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost) // takes an email address. creates user and responds with auth-token. Possibly a log in link
	router.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)       // sets auth token.

	// Assets
	fs := assets.NewStaticImageFS()
	imageRouter := router.PathPrefix("/static").Subrouter()
	imageRouter.PathPrefix("/content/").Handler(http.StripPrefix("/static/content/", fs)).Methods(http.MethodGet)
	imageRouter.Use(imgHeaders)

	// Restricted
	restrictedRouter := router.PathPrefix("/restricted").Subrouter()
	restrictedRouter.Use(authenticate)

	// API
	apiRouter := restrictedRouter.PathPrefix("/api").Subrouter()
	apiRouter.Use(jsHeaders)

	// Country
	countryRepositoryDb := domain.NewCountryRepositoryDb(db)
	countryHandler := CountryHandler{Service: service.NewCountryService(countryRepositoryDb)}
	countryRouter := apiRouter.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countryHandler.FindAllCountries).Methods(http.MethodGet)
	countryRouter.HandleFunc("/", countryHandler.UpdateCountry).Methods(http.MethodPut)
	countryRouter.HandleFunc("/{slug}", countryHandler.FindOneCountry).Methods(http.MethodGet)

	// User
	userRepositoryDb := domain.NewUserRepositoryDb(db)
	userHandler := UserHandler{Service: service.NewUserService(userRepositoryDb)}
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/", userHandler.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/{slug}", userHandler.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.RemoveUser).Methods(http.MethodDelete)

	// Vote
	voteRepositoryDb := domain.NewVoteRepositoryDb(db)
	voteHandler := VoteHandler{Service: service.NewVoteService(voteRepositoryDb)}
	voteRouter := apiRouter.PathPrefix("/vote").Subrouter()
	voteRouter.HandleFunc("/", voteHandler.CreateVote).Methods(http.MethodPost)
	voteRouter.HandleFunc("/", voteHandler.UpdateVote).Methods(http.MethodPut)
	// voteRouter.HandleFunc("/user/{userId}", voteHandler.VoteByUser).Methods(http.MethodGet)
	// voteRouter.HandleFunc("/country/{countryId}", voteHandler.VoteByUser).Methods(http.MethodGet)

	// Chatroom
	commentRepositoryDb := domain.NewCommentRepositoryDb(db)
	commentService := service.NewCommentService(commentRepositoryDb)
	chatRoomHandler := ChatRoomHandler{
		AuthService:    service.NewAuthService(authRepositoryMem),
		RoomService:    service.NewRoom(commentService),
		CommentService: commentService,
	}
	go chatRoomHandler.RoomService.Run()
	chatRouter := router.PathPrefix("/chat").Subrouter()
	chatRouter.HandleFunc("/connect", chatRoomHandler.Connect)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", appConf.Server.Url, appConf.Server.Port),
		Handler: handlers.CORS(headersOk, originsOk, methodsOk)(router),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server listening on port %s", appConf.Server.Port)
	log.Fatal(server.ListenAndServe())
}
