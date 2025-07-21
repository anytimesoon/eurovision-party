package main

import (
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api"
	"github.com/anytimesoon/eurovision-party/pkg/data"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/gorilla/handlers"
	"github.com/timshannon/bolthold"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var authService service.AuthService

func StartServer(store *bolthold.Store) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logging)

	// Repos
	authRepository := data.NewAuthRepositoryDB(store)
	sessionRepository := data.NewSessionRepositoryDb(store)
	userRepository := data.NewUserRepositoryDb(store)
	commentRepository := data.NewCommentRepositoryDb(store)
	countryRepository := data.NewCountryRepositoryDb(store)
	voteRepository := data.NewVoteRepositoryDb(store)

	// Authentication
	authService = service.NewAuthService(authRepository, sessionRepository, userRepository, conf.App.Secret)
	authHandler := api.AuthHandler{Service: authService}
	router.HandleFunc("/api/login", authHandler.Login).Methods(http.MethodPost) // sets auth token.

	// Assets
	imageHandler := api.AssetHandler{Service: service.NewAssetService()}

	imageRouter := router.PathPrefix("/content").Subrouter()
	imageRouter.PathPrefix("/static/").Handler(http.StripPrefix("/content/static/", api.DefaultAssetHandler())).Methods(http.MethodGet)
	imageRouter.Use(ImgHeaders)

	restrictedImageRouter := imageRouter.PathPrefix("/user").Subrouter()
	restrictedImageRouter.HandleFunc("/avatar/{file}", imageHandler.GetAvatar).Methods(http.MethodGet)
	restrictedImageRouter.HandleFunc("/chat/{file}", imageHandler.GetChatImage).Methods(http.MethodGet)
	restrictedImageRouter.Use(Authenticate)

	restrictedImageUploadRouter := router.PathPrefix("/content/user").Subrouter()
	restrictedImageUploadRouter.HandleFunc("/chat", imageHandler.CreateChatImage).Methods(http.MethodPost)
	restrictedImageUploadRouter.Use(Authenticate, JsHeaders)

	// Restricted
	restrictedRouter := router.PathPrefix("/restricted").Subrouter()
	restrictedRouter.Use(Authenticate)

	// API
	apiRouter := restrictedRouter.PathPrefix("/api").Subrouter()
	apiRouter.Use(JsHeaders)

	// Chatroom
	commentService := service.NewCommentService(commentRepository)
	chatRoomHandler := api.ChatRoomHandler{
		RoomService:    service.NewRoom(commentService),
		CommentService: commentService,
		AuthService:    authService,
	}
	go chatRoomHandler.RoomService.Run()
	chatRouter := router.PathPrefix("/chat").Subrouter()
	chatRouter.HandleFunc("/connect/{token}", chatRoomHandler.Connect)

	// Country
	countryHandler := api.CountryHandler{Service: service.NewCountryService(countryRepository)}
	countryRouter := apiRouter.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countryHandler.GetAllCountries).Methods(http.MethodGet) // admin only
	countryRouter.HandleFunc("/", countryHandler.UpdateCountry).Methods(http.MethodPut)   // admin only
	countryRouter.HandleFunc("/participating", countryHandler.GetParticipatingCountries).Methods(http.MethodGet)
	countryRouter.HandleFunc("/{slug}", countryHandler.GetOneCountry).Methods(http.MethodGet)

	// User
	userHandler := api.UserHandler{
		UserService: service.NewUserService(
			userRepository,
			chatRoomHandler.RoomService.BroadcastUpdate,
			authRepository,
			commentRepository,
			voteRepository),
		AssetService: service.NewAssetService(),
	}
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/", userHandler.UpdateUser).Methods(http.MethodPut)       // admin or current user
	userRouter.HandleFunc("/image", userHandler.UpdateImage).Methods(http.MethodPut) // current user only
	userRouter.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	userRouter.HandleFunc("/registered/{userId}", userHandler.GetRegisteredUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.GetOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.DeleteUser).Methods(http.MethodDelete) // admin only

	// Vote
	voteHandler := api.VoteHandler{Service: service.NewVoteService(voteRepository, chatRoomHandler.RoomService.BroadcastUpdate, commentRepository)}
	voteRouter := apiRouter.PathPrefix("/vote").Subrouter()
	voteRouter.HandleFunc("/", voteHandler.UpdateVote).Methods(http.MethodPut) // current user only
	voteRouter.HandleFunc("/results", voteHandler.GetResults).Methods(http.MethodGet)
	voteRouter.HandleFunc("/results/{userId}", voteHandler.GetResultsByUser).Methods(http.MethodGet)
	voteRouter.HandleFunc("/countryanduser/{slug}", voteHandler.GetVoteByUserAndCountry).Methods(http.MethodGet) // current user only

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{conf.App.HttpProto + conf.App.Domain})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credentials := handlers.AllowCredentials()

	server := &http.Server{
		Addr:    conf.App.ServHost,
		Handler: handlers.CORS(headersOk, originsOk, methodsOk, credentials)(router),
		// Good practice to set timeouts to avoid Slow loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server listening on %s", conf.App.ServHost)
	log.Fatal(server.ListenAndServe())
}
