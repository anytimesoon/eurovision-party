package router

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var authService service.AuthService

func StartServer(db *sqlx.DB) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logging)

	// Authentication
	authRepositoryMem := domain.NewAuthRepositoryDB(db)
	authService = service.NewAuthService(authRepositoryMem)
	authHandler := AuthHandler{Service: authService}
	router.HandleFunc("/api/login", authHandler.Login).Methods(http.MethodPost) // sets auth token.

	// Assets
	imageHandler := AssetHandler{Service: service.NewAssetService()}

	imageRouter := router.PathPrefix("/content").Subrouter()
	imageRouter.PathPrefix("/static/").Handler(http.StripPrefix("/content/static/", assetHandler())).Methods(http.MethodGet)
	imageRouter.Use(imgHeaders)

	restrictedImageRouter := imageRouter.PathPrefix("/user").Subrouter()
	restrictedImageRouter.HandleFunc("/avatar/{file}", imageHandler.GetAvatar).Methods(http.MethodGet)
	restrictedImageRouter.HandleFunc("/chat/{file}", imageHandler.GetChatImage).Methods(http.MethodGet)
	restrictedImageRouter.Use(authenticate)

	restrictedImageUploadRouter := router.PathPrefix("/content/user").Subrouter()
	restrictedImageUploadRouter.HandleFunc("/chat", imageHandler.CreateChatImage).Methods(http.MethodPost)
	restrictedImageUploadRouter.Use(authenticate, jsHeaders)

	// Restricted
	restrictedRouter := router.PathPrefix("/restricted").Subrouter()
	restrictedRouter.Use(authenticate)

	// API
	apiRouter := restrictedRouter.PathPrefix("/api").Subrouter()
	apiRouter.Use(jsHeaders)

	// Chatroom
	commentRepositoryDb := domain.NewCommentRepositoryDb(db)
	commentService := service.NewCommentService(commentRepositoryDb)
	chatRoomHandler := ChatRoomHandler{
		RoomService:    service.NewRoom(commentService),
		CommentService: commentService,
		AuthService:    authService,
	}
	go chatRoomHandler.RoomService.Run()
	chatRouter := router.PathPrefix("/chat").Subrouter()
	chatRouter.HandleFunc("/connect/{t}/{u}", chatRoomHandler.Connect)

	// Country
	countryRepositoryDb := domain.NewCountryRepositoryDb(db)
	countryHandler := CountryHandler{Service: service.NewCountryService(countryRepositoryDb)}
	countryRouter := apiRouter.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countryHandler.FindAllCountries).Methods(http.MethodGet) // admin only
	countryRouter.HandleFunc("/", countryHandler.UpdateCountry).Methods(http.MethodPut)    // admin only
	countryRouter.HandleFunc("/participating", countryHandler.Participating).Methods(http.MethodGet)
	countryRouter.HandleFunc("/{slug}", countryHandler.FindOneCountry).Methods(http.MethodGet)

	// User
	userRepositoryDb := domain.NewUserRepositoryDb(db)
	userHandler := UserHandler{Service: service.NewUserService(userRepositoryDb, chatRoomHandler.RoomService.BroadcastUpdate)}
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.FindAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/", userHandler.UpdateUser).Methods(http.MethodPut)        // admin or current user
	userRouter.HandleFunc("/image", userHandler.UpdateImage).Methods(http.MethodPut)  // current user only
	userRouter.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost) // admin only
	userRouter.HandleFunc("/registered", userHandler.FindRegisteredUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.FindOneUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}", userHandler.RemoveUser).Methods(http.MethodDelete) // admin only

	// Vote
	voteRepositoryDb := domain.NewVoteRepositoryDb(db)
	voteHandler := VoteHandler{Service: service.NewVoteService(voteRepositoryDb)}
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
		Addr:    fmt.Sprintf("%s:%s", conf.App.ServHost, conf.App.ServPort),
		Handler: handlers.CORS(headersOk, originsOk, methodsOk, credentials)(router),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server listening on port %s", conf.App.ServPort)
	log.Fatal(server.ListenAndServe())
}
