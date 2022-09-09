package conf

import "net/smtp"

type App struct {
	Server Server
	Auth   Auth
	Email  Email
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

func Setup() App {
	var server = Server{
		Port: "8080",
		Url:  "localhost",
	}

	var auth = Auth{
		SessionKey: "testing-key-session",
	}

	var email = Email{
		UseSSL: false,
	}

	return App{
		Server: server,
		Auth:   auth,
		Email:  email,
	}
}
