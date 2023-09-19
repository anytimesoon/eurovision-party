package conf

import (
	"github.com/google/uuid"
	"net/smtp"
)

type App struct {
	DB      DB
	Server  Server
	Auth    Auth
	Email   Email
	BotUser BotUser
}

type DB struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
}

type Server struct {
	Port string
	Url  string
}

type Auth struct {
	CookieKey  []byte
	SessionKey string
}

type Email struct {
	UseSSL        bool
	Auth          smtp.Auth
	ServerAndPort string
	EmailAddress  string
}

type BotUser struct {
	ID   uuid.UUID
	Name string
}

func Setup() App {
	var db = DB{
		Username: "eurovision",
		Password: "P,PO)+{l4!C{ff",
		Hostname: "127.0.0.1",
		Port:     "3306",
		DBName:   "eurovision",
	}
	var server = Server{
		Port: "8080",
		Url:  "127.0.0.1",
	}

	var auth = Auth{
		SessionKey: "testing-key-session",
	}

	var email = Email{
		UseSSL: false,
	}

	return App{
		DB:     db,
		Server: server,
		Auth:   auth,
		Email:  email,
	}
}
